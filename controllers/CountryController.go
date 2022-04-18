package controllers

import (
	"awesomeProject/controllers/Requests"
	"awesomeProject/controllers/Responses"
	"awesomeProject/models/Countries"
	"awesomeProject/services"
	"encoding/json"
	"github.com/husobee/vestigo"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetCountryByCode(w http.ResponseWriter, r *http.Request) {
	var response, emptyResponse = Countries.GetCountriesByCountryCode(r.FormValue("country_code"))

	var js, _ = json.Marshal(response)

	if emptyResponse != nil {
		w.WriteHeader(http.StatusNotFound)
		js, _ = json.Marshal(emptyResponse)
	}
	_, _ = w.Write(js)
}

func GetCountriesByQuery(w http.ResponseWriter, r *http.Request) {
	var req Requests.RequestQuery
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &req)
	var response, emptyResponse = services.GetCountriesByQuery(req)
	var js, _ = json.Marshal(response)
	if emptyResponse != nil {
		w.WriteHeader(http.StatusNotFound)
		js, _ = json.Marshal(emptyResponse)
	}
	_, _ = w.Write(js)
}

func UpdateCountry(w http.ResponseWriter, r *http.Request) {
	id := vestigo.Param(r, "id")
	var isSupported, _ = strconv.ParseBool(r.FormValue("is_supported"))
	var affectedResults = Countries.UpdateCountryIsSupported(id, isSupported)
	var message = "Country is not updated"

	if affectedResults == 1 {
		message = "Country is updated"
	}

	var js, _ = json.Marshal(Responses.UpdatedResponse{
		Message: message,
	})

	_, _ = w.Write(js)
}
