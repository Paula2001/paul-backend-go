package routers

import (
	"awesomeProject/controllers"
	"github.com/husobee/vestigo"
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

func SetRoutes(vestigo *vestigo.Router) {

	RouteStruct{
		http.StatusOK,
		"/country",
		vestigo.Get,
		controllers.GetCountryByCode,
		govalidator.MapData{
			"country_code": []string{"required", "between:2,3"},
		},
	}.CreateRoute()

	RouteStruct{
		http.StatusOK,
		"/country-lookup",
		vestigo.Get,
		controllers.GetCountriesByQuery,
		nil,
	}.CreateRoute()

	RouteStruct{
		http.StatusOK,
		"/country/:id",
		vestigo.Patch,
		controllers.GetCountriesByQuery,
		nil,
	}.CreateRoute()
}
