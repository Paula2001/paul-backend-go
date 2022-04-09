package routers

import (
	"awesomeProject/controllers"
	"net/http"
)

func SetRoutes(mux *http.ServeMux) {
	CreateRoute(mux, "/test", "GET", controllers.GetTestimonies)
}
