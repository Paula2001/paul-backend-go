package middlewares

import (
	"encoding/json"
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

func Validate(next http.Handler, rules govalidator.MapData) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		opts := govalidator.Options{
			Request:         r,
			Rules:           rules,
			RequiredDefault: true,
		}
		v := govalidator.New(opts)
		e := v.Validate()
		var lenOfErrors = len(e)
		if lenOfErrors != 0 {
			errs := map[string]interface{}{"validationError": e}
			x, _ := json.Marshal(errs)
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write(x)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
