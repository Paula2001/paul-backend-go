package main

import (
	"awesomeProject/Env"
	"awesomeProject/routers"
	"github.com/husobee/vestigo"
	"log"
	"net/http"
	"os"
)

func preStartLogs(port string) {
	log.Println("Listening on :" + port)
}

func main() {
	Env.LoadEnv()
	var port = os.Getenv("APP_PORT")
	router := vestigo.NewRouter()
	preStartLogs(port)
	routers.SetRoutes(router)
	err := http.ListenAndServe(":"+port, router)
	log.Fatal(err)
}
