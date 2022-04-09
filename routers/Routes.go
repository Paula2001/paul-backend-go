package routers

import (
	"awesomeProject/controllers"
	"net/http"
)

func SetRoutes(mux *http.ServeMux) {
	RouteStruct{
		http.StatusOK,
		"/testimone",
		"GET",
		mux,
		controllers.GetTestimonies,
	}.CreateRoute()

	RouteStruct{
		http.StatusOK,
		"/testimone-create",
		"POST",
		mux,
		controllers.GetTestimonies,
	}.CreateRoute()
}
