package utils

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"time"
)

func FordardGet(h string, timeout int, r *http.Request, w http.ResponseWriter) string {
	req, err := http.NewRequest("GET", h, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return ""
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

    TraceID := r.Header.Get("X-B3-Traceid")

   	// create a new header object
	header := http.Header{}

	// add a custom header to the header object
	header.Set("X-B3-Traceid", TraceID)

	// set the header object to the request's Header field
	req.Header = header

	// set timeout
	client := &http.Client{
		Timeout:   time.Second * time.Duration(timeout),
		Transport: tr,
	}

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return "[ERROR] error occur when request"
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return "[ERROR] error occur when getbody"
	}
	defer resp.Body.Close()
	return string(body)
}
