package middlewares

import "net/http"

func SetDefaultHeaders(next http.Handler, statusCode int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})

}
