package main

import (
	"golang-microsvc/models"
	"golang-microsvc/router"
	"log"
	"net/http"
)

func main() {
	defer models.DisconnectMongo()

	r := router.ConfigRouter()
	
	log.Println("Our server is started")
	http.ListenAndServe(":3000", r)
}

