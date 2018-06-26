package main

import (
	"log"
	"net/http"
	"super/controllers"
	"super/system"
)

func main() {
	log.Print("Running")

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Make a webpage
	http.HandleFunc("/", controllers.HomeIndex)
	http.HandleFunc("/users", controllers.UsersIndex)

	log.Fatal(http.ListenAndServe(":8080", system.LogRequest(http.DefaultServeMux)))

}
