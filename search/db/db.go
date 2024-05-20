package db

import (
	"fmt"
	"search/config"

	logger "github.com/sirupsen/logrus"
	"github.com/stevenferrer/solr-go"

	log "github.com/sirupsen/logrus"
)

type SolrClient struct {
	Client     *solr.JSONClient
	Collection string
}

func NewSolrClient(host string, port int, collection string) *SolrClient {
	logger.Debug(fmt.Sprintf("Connecting to Solr at %s:%d", host, port))
	Client := solr.NewJSONClient(fmt.Sprintf("http://%s:%d", config.SOLRHOST, config.SOLRPORT))

	var err error

	if err != nil {
		log.Info("Failed to connect to Solr")
		log.Fatal(err)
	} else {
		log.Info("Connected to Solr successfully")
	}

	return &SolrClient{
		Client:     Client,
		Collection: collection,
	}
}
