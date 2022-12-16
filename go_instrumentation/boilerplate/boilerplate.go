package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Response struct {
	IP      string `json:"ip"`
	Message string `json:"message"`
}

func main() {
	// Start the application
	startMyApp()
}

func startMyApp() {

	port := os.Getenv("APP_LISTENING_PORT")
	if port == "" {
		port = "8000"
	}

	router := mux.NewRouter()
	router.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {
		res := Response{IP: r.RemoteAddr, Message: "Hello From Go!"}
		json.NewEncoder(rw).Encode(res)
	}).Methods("GET")

	log.Println("Starting the application server...")
	http.ListenAndServe(":"+port, router)
}
