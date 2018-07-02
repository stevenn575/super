package main

import (
	"log"
	"net/http"
	"super/controllers"
	"super/system"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log.Print("Running")

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	r := mux.NewRouter()
	r.HandleFunc("/", controllers.HomeIndex)
	r.HandleFunc("/users", controllers.UsersIndex)
	r.HandleFunc("/users/delete/{id:[0-9]+}", controllers.UsersDelete)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", system.LogRequest(http.DefaultServeMux)))

}
