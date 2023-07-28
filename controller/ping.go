package controller

import (
	"log"
	"net/http"
)

func registerPingRoute() {
	http.HandleFunc("/ping", pingHandler)
}

// health check
func pingHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("received ping.")
}
