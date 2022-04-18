package controllers

import (
	"awesomeProject/controllers/Responses"
	"awesomeProject/models/Countries"
	"encoding/json"
	"github.com/husobee/vestigo"
	"io/ioutil"
	"net/http"
	"strconv"
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
	var isFirst = true // Todo : need to move this to the service
	var values []interface{}
	var keys = "where "
	if req.Query.Id != 0 {
		keys += "id = ?"
		values = append(values, req.Query.Id)
		isFirst = false
	}

	if req.Query.Numcode != 0 {
		if !isFirst {
			keys += " OR "
		}
		keys += "numcode = ? "
		values = append(values, req.Query.Numcode)
		isFirst = false
	}

	if req.Query.Short_name != "" {
		if !isFirst {
			keys += " OR "
		}
		keys += "short_name = ?"
		values = append(values, req.Query.Short_name)
		isFirst = false
	}

	if req.Query.Long_name != "" {
		if !isFirst {
			keys += " OR "
		}
		keys += "long_name = ? "
		values = append(values, req.Query.Long_name)
		isFirst = false
	}

	if req.Query.Iso2 != "" {
		if !isFirst {
			keys += " OR "
		}
		keys += "iso2 = ? "
		values = append(values, req.Query.Iso2)
		isFirst = false
	}

	if req.Query.Iso3 != "" {
		if !isFirst {
			keys += " OR "
		}
		keys += "iso3 = ?"
		values = append(values, req.Query.Iso3)
		isFirst = false
	}
	//log.Fatal(keys)
	var response, emptyResponse = Countries.GetCountriesByQuery(keys, values)
	var js, _ = json.Marshal(response)
	if emptyResponse != nil {
		w.WriteHeader(http.StatusNotFound)
		js, _ = json.Marshal(emptyResponse)
	}
	_, _ = w.Write(js) // Todo : create mapper for the response
}

func UpdateCountry(w http.ResponseWriter, r *http.Request) {
	id := vestigo.Param(r, "id") // url params live in the request
	var isSupported, _ = strconv.ParseBool(r.FormValue("is_supported"))
	var affectedResults = Countries.UpdateCountryIsSupported(id, isSupported)
	var message = "Country is not updated"

	if affectedResults == 1 {
		message = "Country is updated"
	}

	var js, _ = json.Marshal(Responses.UpdatedResponse{
		Message: message,
	})

	_, _ = w.Write(js) // Todo : create mapper for the response
}
