package models

import (
	"os"
	"strconv"
	"testing"
)

func TestMain(m *testing.M) {
	DeleteAllTestimonies()
	code := m.Run()
	os.Exit(code)
}

func TestTestimoniesGet(t *testing.T) {
	TestimoniesCreate{
		ClientName:      "name one",
		ClientTestimony: "Good",
	}.Create()
	testimonies := GetAllTestimonies()
	var testimoniesLen int = len(testimonies)

	if testimoniesLen != 1 {
		var ErrorMessage = "Number of Testimonies not right ,len = " + strconv.Itoa(testimoniesLen)
		t.Error(ErrorMessage)
	}
}
