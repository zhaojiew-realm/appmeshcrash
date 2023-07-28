package controller

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/zhojiew/appmesh-load/model"
	"github.com/zhojiew/appmesh-load/utils"
)

var responseStatus = http.StatusOK
var hostUID = utils.GetHostUID()
var m string = utils.GetMultiHost()
var responseMutex = &sync.Mutex{}
var t = utils.GetEcho()
var loadprecision = utils.GetLoadPrecision()

var fh = model.RemoteHost{
	Host: utils.GetForwardHost(),
	Port: utils.GetForwardPort(),
	Path: utils.GetForwardPath(),
}

var delay, _ = strconv.Atoi(utils.GetDelay())
var timeout, _ = strconv.Atoi(utils.GetTimeout())

func RegisterRoutes() {
	registerForwardRoute()
	registerGetRoute()
	registerPingRoute()
	registerFaultRoute()
	registerMultiRoute()
	registerHelloRoute()
}
