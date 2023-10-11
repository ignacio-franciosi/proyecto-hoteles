package connections

import (
	"strconv"

	"github.com/stevenferrer/solr-go"
)

type SolrClient struct {
	Client     *solr.JSONClient
	Collection string
}

func NewSolrClient(host string, port int, collection string) *SolrClient {
	Client := solr.NewJSONClient("http://" + host + ":" + strconv.Itoa(port))
	return &SolrClient{
		Client:     Client,
		Collection: collection,
	}
}
