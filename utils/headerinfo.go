package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func Infohead(r *http.Request) {
	log.Printf("received %s request from %s \n", r.Method, r.RemoteAddr)
	headersJson, err := json.Marshal(r.Header)
	if err != nil {
		log.Printf("Error marshaling headers to JSON: %s \n", err)
	} else {
		log.Printf("Headers: %s \n", headersJson)
	}
	log.Printf("Body: %v", r.Body)
}
