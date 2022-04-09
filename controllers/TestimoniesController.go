package controllers

import (
	"encoding/json"
	"net/http"
)

type GetTestimonyMapper struct {
	Name    string
	Hobbies []string
}

func GetTestimonies(w http.ResponseWriter, r *http.Request) {
	profile := GetTestimonyMapper{"Alex", []string{"snowboarding", "programming"}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(js)
}

func UpdateTestimonies() {

}
