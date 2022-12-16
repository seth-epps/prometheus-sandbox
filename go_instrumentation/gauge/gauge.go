package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Response struct {
	IP      string `json:"ip"`
	Message string `json:"message"`
}

var REQUEST_INPROGRESS = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "go_app_requests_inprogress",
	Help: "Number of application requests in progress",
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
		REQUEST_INPROGRESS.Inc()
		res := Response{IP: r.RemoteAddr, Message: "Hello From Go!"}
		time.Sleep(5 * time.Second)
		json.NewEncoder(rw).Encode(res)

		REQUEST_INPROGRESS.Dec()
	}).Methods("GET")

	log.Println("Starting the gauge application server...")
	router.Path("/metrics").Handler(promhttp.Handler())
	http.ListenAndServe(":"+port, router)
}
