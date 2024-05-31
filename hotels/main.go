package main

import (
	"hotels/app"
	"hotels/db"
	"hotels/queue"
)

func main() {

	//db.Init_db()
	//queue.QueueProducer.InitQueue()
	app.StartRoute()
}
