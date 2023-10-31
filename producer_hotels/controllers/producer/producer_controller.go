package producer

import (
	"producer/config"
	"producer/services"
	client "producer/services/repositories"
	con "producer/utils/connections"
)

var (
	Producer = services.NewProducer(
		(*client.QueueClient)(con.NewQueueClient(config.RABBITUSER, config.RABBITPASSWORD, config.RABBITHOST, config.RABBITPORT)),
	)
)

func StartProducer() {

	Producer.TopicProducer("*.delete")

}
