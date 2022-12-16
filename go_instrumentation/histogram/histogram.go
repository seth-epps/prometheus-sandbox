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

var REQUEST_RESPOND_TIME = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name: "go_app_response_latency_seconds",
	Help: "Response latency in seconds.",
}, []string{"path"})

func routeMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start_time := time.Now()
		route := mux.CurrentRoute(r)
		path, _ := route.GetPathTemplate()

		next.ServeHTTP(w, r)
		time_taken := time.Since(start_time)
		REQUEST_RESPOND_TIME.WithLabelValues(path).Observe(time_taken.Seconds())

	})

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
		time.Sleep(2 * time.Second)
		json.NewEncoder(rw).Encode(res)

	}).Methods("GET")

	router.Use(routeMiddleware)
	log.Println("Starting the histogram application server...")
	router.Path("/metrics").Handler(promhttp.Handler())
	http.ListenAndServe(":"+port, router)
}
