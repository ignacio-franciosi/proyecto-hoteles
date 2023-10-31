package config

var (
	QUEUENAME = "worker_solr" //consumer_solr
	EXCHANGE  = "items"       //hotels

	LBHOST = "lbbusqueda"
	LBPORT = 80 //a chequear

	RABBITUSER     = "guest"
	RABBITPASSWORD = "guest"
	RABBITHOST     = "localhost"
	RABBITPORT     = 5672
)
