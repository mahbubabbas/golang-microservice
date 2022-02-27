package controllers

import (
	"encoding/json"
	"golang-microsvc/models"
	"golang-microsvc/utils"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/gorilla/mux"
)

type PrintJob struct {
	Format    string `json:"format" binding:"required"`
	InvoiceId int    `json:"invoiceId" binding:"required,gte=0"`
	JobId     int    `json:"jobId" binding:"gte=0"`
}

func PrintStudent(rw http.ResponseWriter, r *http.Request) {
	printClient := resty.New()
	p := PrintJob{}

	resp, err := printClient.R().
		SetBody(PrintJob{Format: "A4", InvoiceId: 12, JobId: 10}).
		SetResult(&p).
		Post("http://localhost:5000/print")

	if err != nil {
		utils.Respond(rw, utils.ErrorResponse("Unable to print"))
		return
	}

	utils.Respond(rw, string(resp.Body()))
}

func GetStudents(rw http.ResponseWriter, r *http.Request) {
	usrInfo := r.Context().Value("user").(models.User)
	log.Println(usrInfo)
	if usrInfo.Role != "admin" {
		utils.Respond(rw, utils.ErrorResponse("User is not authorized to access this resource"))
		return
	}

	s := models.Student{}
	utils.Respond(rw, s.FindStudents())
}

func InsertStudents(rw http.ResponseWriter, r *http.Request) {
	s := models.Student{}
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		utils.Respond(rw, utils.ErrorResponse("Input parsing error"))
		return
	}

	resp := s.Insert(&s)
	utils.Respond(rw, resp)
}

func InsertStudentsMany(rw http.ResponseWriter, r *http.Request) {
	sArr := []models.Student{}
	err := json.NewDecoder(r.Body).Decode(&sArr)
	if err != nil {
		utils.Respond(rw, utils.ErrorResponse("Input parsing error"))
		return
	}
	s := models.Student{}

	resp := s.InsertMany(sArr)
	utils.Respond(rw, resp)
}

func UpdateStudent(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	s := models.Student{}
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		utils.Respond(rw, utils.ErrorResponse("Input decode error"))
		return
	}

	utils.Respond(rw, s.Update(id))
}

func UpdateStudents(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ids := vars["id"]
	sArr := []models.Student{}

	err := json.NewDecoder(r.Body).Decode(&sArr)
	if err != nil {
		utils.Respond(rw, utils.ErrorResponse("Unable to parse"))
		return
	}

	s := models.Student{}
	utils.Respond(rw, s.UpdateMany(sArr, ids))
}

func DeleteStudent(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	s := models.Student{}

	utils.Respond(rw, s.Delete(id))
}
