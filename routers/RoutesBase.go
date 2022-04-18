package routers

import (
	"awesomeProject/middlewares"
	"net/http"
)

func (routeMetaData RouteStruct) CreateRoute() {
	routeMetaData.requestMethod(
		routeMetaData.url,
		routeMetaData.logic,
		func(handlerFunc http.HandlerFunc) http.HandlerFunc {
			return middlewares.Validate(handlerFunc, routeMetaData.validateRules)
		},
	)
}
