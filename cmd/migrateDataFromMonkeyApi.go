package main

import (
	"awesomeProject/cmd/structs"
	"awesomeProject/services"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func callMonkeyGetEndpoint() structs.MonkeyDataResponseStruct {
	response, err := http.Get("https://md-api.monkeydata.com/common/country/list")
	if err != nil {
		println(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var responseObject structs.MonkeyDataResponseStruct
	json.Unmarshal(responseData, &responseObject)

	return responseObject
}

func main() {
	start := time.Now()

	var response = callMonkeyGetEndpoint()
	services.CreateManyCountriesThroughMonkeyAPI(response.Data)
	elapsed := time.Since(start)
	log.Printf("The migration is ended in %s", elapsed)

}
