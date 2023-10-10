package main

import (
	"uba/app"
	"uba/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}
