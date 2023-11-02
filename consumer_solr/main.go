package main

import (
	"consumer_solr/controllers/consumer"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	log.Info("Starting consumer")
	consumer.StartConsumer()
}
