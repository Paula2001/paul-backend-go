package routers

import (
	"awesomeProject/middlewares"
	"net/http"
)

func (routeMetaData RouteStruct) CreateRoute() {
	response := http.HandlerFunc(routeMetaData.logic)
	CheckHTTPVars := middlewares.CheckHTTPVars(response, routeMetaData.requestVar)
	Validation := middlewares.Validate(CheckHTTPVars, routeMetaData.validateRules)
	routeMetaData.mux.Handle(routeMetaData.url, Validation)
}
