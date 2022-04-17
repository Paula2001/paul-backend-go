package controllers

import (
	"awesomeProject/models/Countries"
	"encoding/json"
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

func GetCountryByCode(w http.ResponseWriter, r *http.Request) {
	// Todo : create validation
	rules := govalidator.MapData{
		"country_code": []string{"required", "between:2,3"},
	}

	opts := govalidator.Options{
		Request:         r,
		Rules:           rules,
		RequiredDefault: true,
	}
	w.Header().Set("Content-type", "application/json")

	v := govalidator.New(opts)
	e := v.Validate()
	var lenOfErrors = len(e)
	if lenOfErrors != 0 {
		errr := map[string]interface{}{"validationError": e}
		if errr["validationError"] != nil {
			x, _ := json.Marshal(errr)
			w.WriteHeader(http.StatusBadRequest)
			//fmt.Println(len(errr["validationError"].(interface{})))
			_, _ = w.Write(x)
			return
		}
	}
	var response = Countries.GetCountriesByCountryCode(r.FormValue("country_code"))
	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(js) // Todo : create mapper for the response
}
