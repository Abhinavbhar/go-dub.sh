package main

import (
	"Abhinavbhar/dub.sh/redis"
	"Abhinavbhar/dub.sh/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	//database.Connect()
	redis.InitRedis()
	r := mux.NewRouter()
	r.HandleFunc("/", routes.Home).Methods("GET")
	r.HandleFunc("/", routes.Url).Methods("POST")
	r.HandleFunc("/url/{value}", routes.BaseUrl).Methods("GET")

	log.Println("starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
