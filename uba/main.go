package main

import (
	"uba/app"
	"uba/db"
	"uba/utils"
)

func main() {

	db.StartDbEngine()
	app.StartRoute()
	utils.Init_Cache()
}
