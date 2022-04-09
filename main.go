package main

import (
	"awesomeProject/routers"
	"log"
	"net/http"
)

func preStartLogs() {
	log.Println("Listening on :3000...")
}

func main() {
	mux := http.NewServeMux()
	preStartLogs()
	routers.SetRoutes(mux)
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
