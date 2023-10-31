package services

import (
	"fmt"
	"net/http"
	"producer/config"
	client "producer/services/repositories"

	log "github.com/sirupsen/logrus"
)

type ProducerService struct {
	queue *client.QueueClient
}

func NewProducer(
	queue *client.QueueClient,
) *ProducerService {
	return &ProducerService{
		queue: queue,
	}
}

func (s *ProducerService) TopicProducer(topic string) {
	err := s.queue.ProcessMessages(config.EXCHANGE, topic, func(id string) {
		client := &http.Client{}
		//Revisar http
		req, err := http.NewRequest("DELETE", fmt.Sprintf("http://%s:%d/users/%s/items", config.ITEMSHOST, config.ITEMSPORT, id), nil)
		log.Debug("Item delete sent " + id)

		if err != nil {
			log.Debug("error in delete request")
		}

		_, err = client.Do(req)
		if err != nil {
			log.Error(err)
		}

	})
	if err != nil {
		log.Error("Error starting worker processing", err)
	}
}
