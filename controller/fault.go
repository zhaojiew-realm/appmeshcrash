package controller

import (
	"fmt"
	"log"
	"net/http"
)

func registerFaultRoute() {
	http.HandleFunc("/fault", faultHandler)
	http.HandleFunc("/recover", recoverHandler)
}

func faultHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("========== fault ===========")
	responseMutex.Lock()
	defer responseMutex.Unlock()
	responseStatus = http.StatusInternalServerError
	log.Println("received fault request, now returning status ", responseStatus)
	fmt.Fprintf(w, "host: %s will now respond with %d on /get.\n", hostUID, responseStatus)
}

func recoverHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("========== recover ===========")
	responseMutex.Lock()
	defer responseMutex.Unlock()
	responseStatus = http.StatusOK
	log.Println("received recover request, now returning status ", responseStatus)
	fmt.Fprintf(w, "host: %s will now respond with %d on /get. \n", hostUID, responseStatus)
}
