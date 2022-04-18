package routers

import (
	"github.com/husobee/vestigo"
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

type RouteStruct struct {
	status        int
	url           string
	requestMethod func(path string, handler http.HandlerFunc, middleware ...vestigo.Middleware)
	logic         func(w http.ResponseWriter, r *http.Request)
	validateRules govalidator.MapData
}
