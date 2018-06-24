package main

import (
	"log"
	"net/http"
	"super/controllers"
)

func main() {
	log.Print("Running")

	// Make a webpage
	http.HandleFunc("/", controllers.HomeIndex)
	http.HandleFunc("/users", controllers.UsersIndex)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
