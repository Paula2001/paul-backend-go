package routers

import "net/http"

type RouteStruct struct {
	status     int
	url        string
	requestVar string
	mux        *http.ServeMux
	logic      func(w http.ResponseWriter, r *http.Request)
}
