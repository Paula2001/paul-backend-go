package Countries

import (
	"awesomeProject/database"
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

func GetCountriesByCountryCode(countryCode string) []CountryStruct {
	var query = "select * from countries where iso2 = ? OR iso3 = ?"
	results, err := database.Connection.Query(query, countryCode, countryCode)
	if err != nil {
		panic(err.Error())
	}
	var countries []CountryStruct
	countries = []CountryStruct{}
	for results.Next() {
		var country CountryStruct
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
			println("this is error", err.Error())
			return nil
		}
		countries = append(countries, country)
	}
	return countries
}
