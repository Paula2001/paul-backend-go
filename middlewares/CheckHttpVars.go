package middlewares

import (
	"net/http"
)

func CheckHTTPVars(next http.Handler, requestVar string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if requestVar != r.Method {
			var errorMSG = "Error no match route for this method " + requestVar + " and for this route " + r.RequestURI
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errorMSG))
		} else {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		}
	})
}
