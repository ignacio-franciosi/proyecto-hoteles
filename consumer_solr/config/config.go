package config

var (
	QUEUENAME = "consumer_solr" //worker_solr
	EXCHANGE  = "hotels"        //items

	LBHOST = "lbbusqueda"
	LBPORT = 80 //a chequear

	RABBITUSER     = "guest"
	RABBITPASSWORD = "guest"
	RABBITHOST     = "localhost"
	RABBITPORT     = 5672
)
