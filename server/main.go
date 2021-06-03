package main

import (
	"fmt"
	"log"
	"net/http"

	"server/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on 8080...")
	buildHandler := http.FileServer(http.Dir("/Users/Apple/GoogleMatt/Web Programming/react-golang-mongodb-todolist/build"))
	r.PathPrefix("/").Handler(buildHandler)
	staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("/Users/Apple/GoogleMatt/Web Programming/react-golang-mongodb-todolist/build/static")))
	r.PathPrefix("/static/").Handler(staticHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
