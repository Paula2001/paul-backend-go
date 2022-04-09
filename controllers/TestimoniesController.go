package controllers

import (
	"awesomeProject/models"
	"encoding/json"
	"net/http"
)

func GetTestimonies(w http.ResponseWriter, r *http.Request) {
	response := models.GetAllTestimonies()

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(js)
}

func CreateTestimonies(w http.ResponseWriter, r *http.Request) {
	response := createTestimonyResponseStruct{Message: "Testimony created successfully"}
	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	models.TestimoniesCreate{
		ClientName:      r.FormValue("client-name"),
		ClientTestimony: r.FormValue("client-testimony"),
	}.Create()

	_, _ = w.Write(js) // Todo : should return the created data
}
