package controllers

import (
	"awesomeProject/models/Countries"
	"encoding/json"
	"net/http"
)

func GetCountryByCode(w http.ResponseWriter, r *http.Request) {
	var response, emptyResponse = Countries.GetCountriesByCountryCode(r.FormValue("country_code"))

	var js, _ = json.Marshal(response)

	if emptyResponse != nil {
		w.WriteHeader(http.StatusNotFound)
		js, _ = json.Marshal(emptyResponse)
	}
	_, _ = w.Write(js) // Todo : create mapper for the response
}
