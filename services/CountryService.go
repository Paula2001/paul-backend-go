package services

import (
	"awesomeProject/cmd/structs"
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
