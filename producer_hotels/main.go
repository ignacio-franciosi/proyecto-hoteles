package main

import (
	producer "producer/controllers/producer"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	log.Info("Starting producer")
	producer.StartWorker()
}
