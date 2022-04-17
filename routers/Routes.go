package routers

import (
	"awesomeProject/controllers"
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

func SetRoutes(mux *http.ServeMux) {
	RouteStruct{
		http.StatusOK,
		"/testimone", // Todo : should be able to create routes with the same name
		"GET",        // Todo : refactor
		mux,
		controllers.GetTestimonies,
		govalidator.MapData{
			"country_code": []string{"required", "between:2,3"},
		},
	}.CreateRoute()

	RouteStruct{
		http.StatusCreated,
		"/testimone-create",
		"POST",
		mux,
		controllers.CreateTestimonies,
		govalidator.MapData{},
	}.CreateRoute()

	RouteStruct{
		http.StatusOK,
		"/countries",
		"GET",
		mux,
		controllers.GetCountryByCode,
		govalidator.MapData{
			"country_code": []string{"required", "between:2,3"},
		},
	}.CreateRoute()
}
