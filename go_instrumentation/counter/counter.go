package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Response struct {
	IP      string `json:"ip"`
	Message string `json:"message"`
}

var REQUEST_COUNT = promauto.NewCounter(prometheus.CounterOpts{
	Name: "go_app_requests_count",
	Help: "Total App HTTP Requests count.",
})

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

		REQUEST_COUNT.Inc()
	}).Methods("GET")

	log.Println("Starting the counter application server...")
	router.Path("/metrics").Handler(promhttp.Handler())
	http.ListenAndServe(":"+port, router)
}
