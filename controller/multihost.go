package controller

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/zhaojiew-realm/appmeshcrash/utils"
)

func registerMultiRoute() {
	http.HandleFunc("/multihost", multihostHandler)
}

func multihostHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("========== multihost ===========")

	// output request info
	utils.Infohead(r)

	// load
	utils.CpuLoad(loadprecision)

	// set response delay
	log.Printf("delay for %d second \n", delay)
	time.Sleep(time.Duration(delay) * time.Second)

	// get request id
	requestID := r.Header.Get("X-Request-Id")

	// fordard to multi host
	ms := strings.Split(m, ",")
	for _, h := range ms {
		log.Printf("starting forward to %s \n", h)

		//send back customer header with hostUID
		w.Header().Add("HostUID", hostUID)
		w.WriteHeader(responseStatus)
		if responseStatus == http.StatusOK {
			// forward request to upstream server
			resp := utils.ForwardGet(h, timeout, r, w, r.Header)
			log.Printf("successfully forward request to upstream server %s \n", h)
			fmt.Fprintf(w, "[MULTI] %s (%s)==> %s \n", t, requestID, resp)
		} else {
			fmt.Fprintf(w, "%s no response 4 u", hostUID)
		}
	}

}
