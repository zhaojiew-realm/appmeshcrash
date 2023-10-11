package controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/zhaojiew-realm/appmeshcrash/utils"
)

func registerGetRoute() {
	http.HandleFunc("/get", getHandler)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("========== get ===========")

	// output request info
	utils.Infohead(r)

	// load
	utils.CpuLoad(loadprecision)

	// set response delay
	log.Printf("delay for %d second \n", delay)
	time.Sleep(time.Duration(delay) * time.Second)

	// send back customer header with hostUID
	w.Header().Add("HostUID", hostUID)
	w.WriteHeader(responseStatus)

	if responseStatus == http.StatusOK {
		fmt.Fprintf(w, "[GET] "+t)
	} else {
		fmt.Fprintf(w, "%s no response 4 u", hostUID)
	}
}
