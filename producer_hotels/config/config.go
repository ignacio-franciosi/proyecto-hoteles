package config

var (
	QUEUENAME = "producer_hotels" //worker_items
	EXCHANGE  = "hotels"          //revisar: users

	ITEMSHOST = "hotels"
	ITEMSPORT = 8090

	RABBITUSER     = "guest"
	RABBITPASSWORD = "guest"
	RABBITHOST     = "localhost"
	RABBITPORT     = 5672
)
