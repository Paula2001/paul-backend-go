package models

import (
	"awesomeProject/database"
	"fmt"
)

type TestimoniesGet struct {
	ClientName      string
	ClientTestimony string
}
type TestimoniesCreate struct {
	ClientName      string
	ClientTestimony string
}

func (t TestimoniesCreate) Create() {
	preparedStat, _ := database.Connection.Prepare("insert into testimonies(name, testimony) values(?,?)")
	preparedStat.Exec(t.ClientName, t.ClientTestimony)
}

func GetAllTestimonies() []TestimoniesGet {
	results, err := database.Connection.Query("select name, testimony from testimonies")
	if err != nil {
		panic(err.Error())
	}
	var testimonies []TestimoniesGet
	for results.Next() {
		var testimony TestimoniesGet
		err = results.Scan(&testimony.ClientName, &testimony.ClientTestimony)
		if err != nil {
			panic(err.Error())
		}
		testimonies = append(testimonies, testimony)
	}
	fmt.Println(testimonies[0].ClientName)
	return testimonies
}

func DeleteAllTestimonies() {
	database.Connection.Query("delete from testimonies")
}
