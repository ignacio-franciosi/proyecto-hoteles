package consumer

import (
	"consumer_solr/config"
	"consumer_solr/services"
	client "consumer_solr/services/repositories"
	con "consumer_solr/utils/connections"
)

var (
	Consumer = services.NewConsumer(
		(*client.QueueClient)(con.NewQueueClient(config.RABBITUSER, config.RABBITPASSWORD, config.RABBITHOST, config.RABBITPORT)),
	)
)

func StartConsumer() {

	Consumer.TopicConsumer("*.*")

}
