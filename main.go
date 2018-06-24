package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"super/controllers"
	"time"
)

func main() {
	log.Print("Running")

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Make a webpage
	http.HandleFunc("/", controllers.HomeIndex)
	http.HandleFunc("/users", controllers.UsersIndex)

	log.Fatal(http.ListenAndServe(":8080", logRequest(http.DefaultServeMux)))

}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Looks like 2018/06/24 10:45:43 [::1]:49731 GET /
		logLine := fmt.Sprintf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

		// Log to console
		log.Printf(logLine)

		// Write to log
		go writeLog(logLine)

		handler.ServeHTTP(w, r)
	})
}

func writeLog(line string) {
	// Simulates delay because this is asyncronous
	time.Sleep(2 * time.Second)

	f, err := os.OpenFile("./logs/application.log", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString(line); err != nil {
		panic(err)
	}
}
