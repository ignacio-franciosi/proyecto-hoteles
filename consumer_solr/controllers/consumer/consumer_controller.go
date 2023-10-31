package consumer

import (
	"consumer/config"
	"consumer/services"
	client "consumer/services/repositories"
	con "consumer/utils/connections"
)

var (
	Consumer = services.NewConsumer(
		(*client.QueueClient)(con.NewQueueClient(config.RABBITUSER, config.RABBITPASSWORD, config.RABBITHOST, config.RABBITPORT)),
	)
)

func StartConsumer() {

	Consumer.TopicConsumer("*.*")

}
