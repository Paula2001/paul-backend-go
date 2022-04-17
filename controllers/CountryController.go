package controllers

import (
	"awesomeProject/models/Countries"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type QueryObject struct {
	Id         int
	Short_name string
	Long_name  string
	Numcode    int
	Iso2       string
	Iso3       string
}

type requestQuery struct {
	Query QueryObject
}

func GetCountryByCode(w http.ResponseWriter, r *http.Request) {
	var response, emptyResponse = Countries.GetCountriesByCountryCode(r.FormValue("country_code"))

	var js, _ = json.Marshal(response)

	if emptyResponse != nil {
		w.WriteHeader(http.StatusNotFound)
		js, _ = json.Marshal(emptyResponse)
	}
	_, _ = w.Write(js) // Todo : create mapper for the response
}

func GetCountriesByQuery(w http.ResponseWriter, r *http.Request) {
	var req requestQuery
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &req)
	var values []interface{}
	var keys = "where "
	if req.Query.Id != 0 {
		keys += "id = ? OR "
		values = append(values, req.Query.Id)
	}

	if req.Query.Numcode != 0 {
		keys += "numcode = ? OR "
		values = append(values, req.Query.Numcode)
	}

	if req.Query.Short_name != "" {
		keys += "short_name = ? OR "
		values = append(values, req.Query.Short_name)
	}

	if req.Query.Long_name != "" {
		keys += "long_name = ? OR "
		values = append(values, req.Query.Long_name)
	}

	if req.Query.Iso2 != "" {
		keys += "iso2 = ? OR "
		values = append(values, req.Query.Iso2)
	}

	if req.Query.Iso3 != "" {
		keys += "iso3 = ?"
		values = append(values, req.Query.Iso3)
	}
	var response, emptyResponse = Countries.GetCountriesByQuery(keys, values)
	var js, _ = json.Marshal(response)
	if emptyResponse != nil {
		w.WriteHeader(http.StatusNotFound)
		js, _ = json.Marshal(emptyResponse)
	}
	_, _ = w.Write(js) // Todo : create mapper for the response
}
