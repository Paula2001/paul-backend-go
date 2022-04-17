package controllers

import (
	"awesomeProject/models/Countries"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	Countries.TruncateCountries()
	code := m.Run()
	os.Exit(code)
}

func createCzechRepublic() {
	Countries.CountryStruct{
		Id:           1,
		Iso2:         "CZ",
		Short_name:   "Czech Republic",
		Long_name:    "Czech Republic",
		Iso3:         "CZE",
		Numcode:      203,
		Is_supported: true,
	}.Create()
}

func createUsa() {
	Countries.CountryStruct{
		Id:           2,
		Iso2:         "US",
		Short_name:   "United States",
		Long_name:    "United States Of America",
		Iso3:         "USA",
		Numcode:      203,
		Is_supported: true,
	}.Create()

}

func TestGetCountriesByCode(t *testing.T) {
	t.Run("should return not found by calling /countries?country_code", func(t *testing.T) {
		createCzechRepublic()
		request, _ := http.NewRequest(http.MethodGet, "countries?country_code=BB", nil)
		response := httptest.NewRecorder()

		GetCountryByCode(response, request)

		got := response.Body.String()
		want := "{\"Message\":\"No Country is not found with country code BB\"}"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("should return czech republic when provide it's code by calling /countries?country_code", func(t *testing.T) {
		createCzechRepublic()
		createUsa()
		request, _ := http.NewRequest(http.MethodGet, "countries?country_code=CZ", nil)
		response := httptest.NewRecorder()

		GetCountryByCode(response, request)

		got := response.Body.String()
		want := "{\"Id\":1,\"Iso2\":\"CZ\",\"Short_name\":\"Czech Republic\",\"Long_name\":\"Czech Republic\",\"Iso3\":\"CZE\",\"Numcode\":203,\"Is_supported\":true}"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("should return czech republic and usa by calling /country-look when provide it's code", func(t *testing.T) {
		//createCzechRepublic()
		createUsa()
		m, b := map[string]map[string]interface{}{"query": {"id": 1, "iso2": "US"}}, new(bytes.Buffer)
		json.NewEncoder(b).Encode(m)
		request, _ := http.NewRequest(http.MethodGet, "country-lookup?query=CZ", b)
		response := httptest.NewRecorder()

		GetCountriesByQuery(response, request)

		got := response.Body.String()
		want := "[{\"Id\":1,\"Iso2\":\"CZ\",\"Short_name\":\"Czech Republic\",\"Long_name\":\"Czech Republic\",\"Iso3\":\"CZE\",\"Numcode\":203,\"Is_supported\":true}]"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
