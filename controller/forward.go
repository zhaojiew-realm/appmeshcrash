package controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/zhojiew/appmesh-load/utils"
)

func registerForwardRoute() {
	http.HandleFunc("/forward", forwardHandler)
}

func forwardHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("========== forward ===========")
	utils.Infohead(r)

	h := fh.GetUrl()
	log.Printf("starting forward to %s \n", h)

	// set response delay
	log.Printf("delay for %d second \n", delay)
	time.Sleep(time.Duration(delay) * time.Second)

	// load
	utils.CpuLoad(loadprecision)

	//send back customer header with hostUID
	w.Header().Add("HostUID", hostUID)
	w.WriteHeader(responseStatus)
	if responseStatus == http.StatusOK {
		// forward request to upstream server
		resp := utils.FordardGet(h, timeout, r, w)
		log.Printf("successfully forward request to upstream server %s \n", h)

		fmt.Fprintf(w, "[SINGLE] %s ==> %s \n", t, resp)
	} else {
		fmt.Fprintf(w, "%s no response 4 u", hostUID)
	}
}
