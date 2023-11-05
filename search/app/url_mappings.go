package app

import (
	solrController "search/controllers/solr"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	router.GET("/search/:city/:startDate/:endDate", solrController.GetQuery)

	router.GET("/hotels/:id", solrController.Add)

	router.POST("/solr/hotels", solrController.AddHotelsToSolr)

	router.DELETE("/hotels/:id", solrController.Delete)

	log.Info("Finishing mappings configurations")
}
