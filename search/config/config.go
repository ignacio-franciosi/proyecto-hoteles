package config

import (
	"fmt"
)

var (
	SOLRHOST       = "solr"
	SOLRPORT       = 8983
	SOLRCOLLECTION = "hotels"

	HOTELSHOST = "hotels"
	HOTELSPORT = 8090

	QUEUENAME = "consumer_solr" //worker_solr
	EXCHANGE  = "hotels"        //items

	LBHOST = "lbbusqueda"
	LBPORT = 80 //a chequear

	RABBITUSER     = "guest"
	RABBITPASSWORD = "guest"
	RABBITHOST     = "localhost"
	RABBITPORT     = 5672

	AMPQConnectionURL = fmt.Sprintf("amqp://%s:%s@%s:%d/", RABBITUSER, RABBITPASSWORD, RABBITHOST, RABBITPORT)

	USERAPIHOST = "uba" //"user-res-api"
	USERAPIPORT = 8070
)
