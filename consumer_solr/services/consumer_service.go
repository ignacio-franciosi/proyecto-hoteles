package services

import (
	"consumer/config"
	client "consumer/services/repositories"
	"fmt"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

type ConsumerService struct {
	queue *client.QueueClient
}

func NewWorker(
	queue *client.QueueClient,
) *ConsumerService {
	return &ConsumerService{
		queue: queue,
	}
}

func (s *ConsumerService) TopicConsumer(topic string) {
	err := s.queue.ProcessMessages(config.EXCHANGE, topic, func(id string) {
		var resp *http.Response
		var err error
		cli := &http.Client{}
		strs := strings.Split(id, ".")
		if len(strs) < 2 {
			resp, err = http.Get(fmt.Sprintf("http://%s:%d/items/%s", config.LBHOST, config.LBPORT, id))
		} else {
			if strs[1] == "delete" {
				req, err := http.NewRequest("DELETE", fmt.Sprintf("http://%s:%d/items/%s", config.LBHOST, config.LBPORT, strs[0]), nil)
				if err != nil {
					log.Error(err)
				}
				resp, err = cli.Do(req)
				if err != nil {
					log.Error(err)
					log.Debug(resp)
				}
			}
		}
		log.Debug("Item sent " + id)
		if err != nil {
			log.Debug("error in get request")
			log.Debug(resp)
		}
	})
	if err != nil {
		log.Error("Error starting worker processing", err)
	}
}
