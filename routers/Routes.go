package routers

import (
	"awesomeProject/controllers"
	"net/http"
)

func SetRoutes(mux *http.ServeMux) {
	RouteStruct{
		http.StatusOK,
		"/testimone", // Todo : should be able to create routes with the same name
		"GET",        // Todo : refactor
		mux,
		controllers.GetTestimonies,
	}.CreateRoute()

	RouteStruct{
		http.StatusCreated,
		"/testimone-create",
		"POST",
		mux,
		controllers.CreateTestimonies,
	}.CreateRoute()
}
