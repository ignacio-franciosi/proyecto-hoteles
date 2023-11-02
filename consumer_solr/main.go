package main

import (
	"consumer/controllers/consumer"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	log.Info("Starting consumer")
	consumer.StartConsumer()
}
