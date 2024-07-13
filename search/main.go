package main

import (
	"search/app"
	q "search/utils/connections"
)

func main() {
	go q.QueueConnection()
	app.StartRoute()
}
