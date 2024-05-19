package config

import (
	"fmt"
)

var (
	SOLRHOST       = "solr"
	SOLRPORT       = 8983
	SOLRCOLLECTION = "hotelsSearch"

	HOTELSHOST = "hotels"
	HOTELSPORT = 8000

	QUEUENAME = "consumer_solr" //worker_solr
	EXCHANGE  = "hotels"        //items

	LBHOST = "lbbusqueda"
	LBPORT = 80 //a chequear

	RABBITUSER     = "user"     //guest
	RABBITPASSWORD = "password" //guest
	RABBITHOST     = "rabbit"   //localhost
	RABBITPORT     = 5672

	AMPQConnectionURL = fmt.Sprintf("amqp://%s:%s@%s:%d/", RABBITUSER, RABBITPASSWORD, RABBITHOST, RABBITPORT)

	USERAPIHOST = "uba" //"user-res-api"
	//UBAAPIHOST
	USERAPIPORT = 8080
	//UBAAPIPORT
)
