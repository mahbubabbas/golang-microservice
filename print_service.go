package main

import (
	"encoding/json"
	"golang-microsvc/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type PrintJob struct {
	Format    string `json:"format" binding:"required"`
	InvoiceId int    `json:"invoiceId" binding:"required,gte=0"`
	JobId     int    `json:"jobId" binding:"gte=0"`
}

func main() {
	log.Println("Print server is started at port 5000")

	r := mux.NewRouter()
	r.HandleFunc("/print", myPrint).Methods("POST")

	http.ListenAndServe(":5000", r)
}

func myPrint(rw http.ResponseWriter, r *http.Request) {
	log.Println("print service called")
	
	p := PrintJob{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		utils.Respond(rw, utils.ErrorResponse("error decoding"))
		return
	}

	utils.Respond(rw, utils.SuccessResponse("Printing successful"))
}
