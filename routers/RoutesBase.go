package routers

import (
	"awesomeProject/middlewares"
	"net/http"
)

func (routeMetaData RouteStruct) CreateRoute() {
	response := http.HandlerFunc(routeMetaData.logic)
	returnResponse := middlewares.SetDefaultHeaders(response, routeMetaData.status)
	CheckHTTPVars := middlewares.CheckHTTPVars(returnResponse, routeMetaData.requestVar) // Todo : need a refactor it doesn't make sense
	Validation := middlewares.Validate(CheckHTTPVars, routeMetaData.validateRules)
	routeMetaData.mux.Handle(routeMetaData.url, Validation)
}
