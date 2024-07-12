package main

import (
	"search2/app"
	"search2/queue"
	"search2/solr"
	"sync"
)

func main() {

	queue.InitQueue()
	solr.InitSolr()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		queue.Consume()
	}()

	app.StartRoute()

	// Wait for the goroutine to finish
	wg.Wait()
}
