package app

import (
	"docker_containers/controller"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	// Add all methods and its mappings
	router.GET("/services", controller.GetScalableServices)
	router.GET("/info", controller.GetInfo)
	router.GET("/info/:service", controller.GetInfoByService)
	router.POST("/scale/:service", controller.ScaleService)
	router.DELETE("/container/:id", controller.DeleteContainer)

	log.Info("Finishing mappings configurations")
}
