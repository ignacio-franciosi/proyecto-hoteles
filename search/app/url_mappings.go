package app

import (
	solrController "search/controllers/solr"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	router.GET("/search/:city/:startDate/:endDate", solrController.GetQuery)
	router.GET("/hotels/:id", solrController.Add)

	log.Info("Finishing mappings configurations")
}
