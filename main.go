package main

import (
	"awesomeProject/routers"
	"log"
	"net/http"
)

func preStartLogs() {
	log.Println("Listening on :3333...")
}

func main() {
	mux := http.NewServeMux()
	preStartLogs()
	routers.SetRoutes(mux)
	err := http.ListenAndServe(":3333", mux)
	log.Fatal(err)
}
