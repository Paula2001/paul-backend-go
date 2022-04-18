package services

import (
	"awesomeProject/cmd/structs"
	"awesomeProject/controllers/Requests"
	"awesomeProject/models"
	"awesomeProject/models/Countries"
	"strconv"
)

func createPrepareStatUponLenOfCountries(len int) string {
	var stat string
	for i := 0; i < len; i++ {
		if i == len-1 {
			stat += "(?,?,?,?,?,?)"
		} else {
			stat += "(?,?,?,?,?,?),"
		}
	}
	return stat
}

func mapCountryFromMonkeyDataCountriesToCountries(monkeyData []structs.MonkeyDataCountriesResponseStruct) []interface{} {
	var values []interface{}
	for _, value := range monkeyData {
		numCode, _ := strconv.Atoi(value.Numcode)
		values = append(values, value.Iso2, value.Short_name, value.Long_name, value.Iso3, numCode, value.Un_member == "yes")
	}
	return values
}

func CreateManyCountriesThroughMonkeyAPI(response []structs.MonkeyDataCountriesResponseStruct) {
	var lenOfCountries = len(response)
	var statQuestionMarks = createPrepareStatUponLenOfCountries(lenOfCountries)
	var data = mapCountryFromMonkeyDataCountriesToCountries(response)
	Countries.CreateMany(statQuestionMarks, data)
}

func GetCountriesByQuery(req Requests.RequestQuery) (*[]Countries.CountryStruct, *models.EmptyResultStruct) {
	var isFirst = true
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
	var response, emptyResponse = Countries.GetCountriesByQuery(keys, values)

	return response, emptyResponse
}
