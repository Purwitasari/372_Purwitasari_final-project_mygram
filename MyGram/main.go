package main

import (
	"MyGram/database"
	"MyGram/routers"
)

func main() {
	database.StartDB()
	r := routers.StartApp()
	r.Run(":8080")
}
