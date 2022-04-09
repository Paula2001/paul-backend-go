package routers

import (
	"awesomeProject/middlewares"
	"net/http"
)

func CreateRoute(
	mux *http.ServeMux,
	url string,
	requestVar string,
	logic func(w http.ResponseWriter, r *http.Request),
) {
	finalHandler := http.HandlerFunc(logic)
	mux.Handle(url, middlewares.CheckHTTPVars(finalHandler, requestVar))
}
