package main

import (
	"docker_containers/app"
	"docker_containers/client"
)

func main() {

	services := client.GetScalableServices()

	for _, service := range services {

		go client.AutoScale(service)

	}

	app.StartRoute()
}
