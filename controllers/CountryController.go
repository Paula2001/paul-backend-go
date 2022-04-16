package controllers

import (
	"awesomeProject/models"
	"encoding/json"
	"net/http"
)

func GetCountryByCode(w http.ResponseWriter, r *http.Request) {
	response := models.GetAllTestimonies()

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(js)
}
