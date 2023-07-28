package controller

import (
	"fmt"
	"net/http"
)

func registerHelloRoute() {
	http.HandleFunc("/", helloHandler)
}

// health check
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to apppmesh load!")
}
