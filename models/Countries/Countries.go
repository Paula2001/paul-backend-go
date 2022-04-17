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

func GetCountriesByQuery(whereQuery string, data []interface{}) (*[]CountryStruct, *models.EmptyResultStruct) {
	var query = "select * from countries " + whereQuery
	results, err := database.Connection.Query(query, data...)
	if err != nil {
		panic(err.Error())
	}

	var count = 0
	var countries []CountryStruct
	for results.Next() {
		var country CountryStruct
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
		countries = append(countries, country)
	}
	if count == 0 {
		return nil, &models.EmptyResultStruct{Message: "No Countries was found with this query"}
	}

	return &countries, nil
}

func (countryStruct CountryStruct) Create() {
	var query = "insert into countries(long_name, short_name, numcode,iso2, iso3, is_supported) values(?,?,?,?,?,?)"
	preparedStat, _ := database.Connection.Prepare(query)
	preparedStat.Exec(
		countryStruct.Long_name,
		countryStruct.Short_name,
		countryStruct.Numcode,
		countryStruct.Iso2,
		countryStruct.Iso3,
		countryStruct.Is_supported,
	)
}

func TruncateCountries() {
	database.Connection.Query("truncate countries")
}
