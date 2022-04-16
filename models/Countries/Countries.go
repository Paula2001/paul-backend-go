package Countries

import (
	"awesomeProject/database"
	"log"
)

func CreateMany(statQuestionMarks string, countryStruct []interface{}) {
	var query = "insert into countries(iso2,short_name,long_name,iso3,numcode,is_supported) values" + statQuestionMarks
	preparedStat, _ := database.Connection.Prepare(query)
	_, err := preparedStat.Exec(countryStruct...)
	if err != nil {
		log.Fatal(err)
		return
	}
}
