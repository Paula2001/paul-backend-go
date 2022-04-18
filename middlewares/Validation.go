package middlewares

import (
	"encoding/json"
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

func Validate(f http.HandlerFunc, rules govalidator.MapData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if rules == nil {
			f(w, r)
			return
		}
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
			f(w, r)
		}
	}
}
