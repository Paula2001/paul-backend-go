package controllers

import (
	"awesomeProject/models/Countries"
	"encoding/json"
	"net/http"
)

func GetCountryByCode(w http.ResponseWriter, r *http.Request) {
	var response = Countries.GetCountriesByCountryCode(r.FormValue("country_code"))
	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(js) // Todo : create mapper for the response
}
