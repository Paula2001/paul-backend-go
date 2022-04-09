package main

import (
	"awesomeProject/routers"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Profile struct {
	Name    string
	Hobbies []string
}

func main() {
	mux := http.NewServeMux()
	//routers.MetaRequestData{"asdf"}
	routers.CreateRoute(mux, "/test", "GET", foo)
	log.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}

func foo(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Alex", []string{"snowboarding", "programming"}}

	if r.Method == http.MethodGet {
		io.WriteString(w, "This is a get request")
		return
	} else if r.Method == http.MethodPost {
		io.WriteString(w, "This is a post request")
		return
	}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}
