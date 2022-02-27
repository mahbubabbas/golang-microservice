package controllers

import (
	"encoding/json"
	"golang-microsvc/models"
	"golang-microsvc/utils"
	"net/http"
)

func Login(rw http.ResponseWriter, r *http.Request) {
	u := models.User{}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		utils.Respond(rw, utils.ErrorResponse("Parse error"))
		return
	}

	utils.Respond(rw, u.Login(u.Username, u.Password))
}

