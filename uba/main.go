package main

import (
	"uba/app"
	"uba/db"
	utils "uba/utils/cache"
)

func main() {
	utils.InitCache()
	db.StartDbEngine()
	app.StartRoute()

}
