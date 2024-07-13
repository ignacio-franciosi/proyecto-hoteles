package solr

import (
	solr "github.com/rtt/Go-Solr"
	log "github.com/sirupsen/logrus"
)

var SolrClient *solr.Connection

func InitSolr() {

	var err error

	SolrClient, err = solr.Init("my-solr", 8983, "hotelsSearch")
	if err != nil {
		log.Info("Failed to connect to Solr")
		log.Fatal(err)
	} else {
		log.Info("Connected to Solr successfully")
	}
}
