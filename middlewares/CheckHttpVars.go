package middlewares

import (
	"net/http"
)

func CheckHTTPVars(next http.Handler, requestVar string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if requestVar != r.Method {
			w.Write([]byte("ERROR"))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
