package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/zhojiew/appmesh-load/controller"
	"github.com/zhojiew/appmesh-load/utils"
)

var hostUID = utils.GetHostUID()

func main() {
	// confirm the version
	log.Println("welcome to appmeshload v1.0")
	// print configure info from env
	log.Printf("starting server on port %s\n", utils.GetServerPort())
	log.Printf("host unique identifer: %s\n", hostUID)
	log.Printf("delay is: %ss\n", utils.GetDelay())
	log.Printf("load pressure is: %d \n", utils.GetLoadPrecision()/2500000000)
	log.Printf("fix response is: %s", utils.GetEcho())
	log.Printf("single forwatd to %s:%s%s\n", utils.GetForwardHost(), utils.GetForwardPort(), utils.GetForwardPath())
	ms := strings.Split(utils.GetMultiHost(), ",")
	for _, h := range ms {
		log.Printf("multi forward to %s\n", h)
	}
	controller.RegisterRoutes()
	log.Fatal(http.ListenAndServe(":"+utils.GetServerPort(), nil))
}
