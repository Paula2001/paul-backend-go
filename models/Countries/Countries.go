package Countries

import (
	"awesomeProject/database"
	"awesomeProject/models"
	"log"
)

func CreateMany(statQuestionMarks string, countryStruct []interface{}) {
	var query = "insert ignore into countries(iso2,short_name,long_name,iso3,numcode,is_supported) values" + statQuestionMarks
	preparedStat, _ := database.Connection.Prepare(query)
	_, err := preparedStat.Exec(countryStruct...)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func GetCountriesByCountryCode(countryCode string) (*CountryStruct, *models.EmptyResultStruct) {
	var query = "select * from countries where iso2 = ? OR iso3 = ? limit 1"
	results, err := database.Connection.Query(query, countryCode, countryCode)
	if err != nil {
		panic(err.Error())
	}

	var country CountryStruct
	var count = 0

	for results.Next() {
		count++
		err := results.Scan(
			&country.Id,
			&country.Iso2,
			&country.Iso3,
			&country.Numcode,
			&country.Long_name,
			&country.Short_name,
			&country.Is_supported,
		)
		if err != nil {
			panic(err.Error())
		}
	}
	if count == 0 {
		return nil, &models.EmptyResultStruct{Message: "No Country is not found with country code " + countryCode}
	}

	return &country, nil
}
