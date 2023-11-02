package producer

import (
	"producer_hotels/config"
	"producer_hotels/services"
	client "producer_hotels/services/repositories"
	con "producer_hotels/utils/connections"
)

var (
	Producer = services.NewProducer(
		(*client.QueueClient)(con.NewQueueClient(config.RABBITUSER, config.RABBITPASSWORD, config.RABBITHOST, config.RABBITPORT)),
	)
)

func StartProducer() {

	Producer.TopicProducer("*.*", "-", "")

}
