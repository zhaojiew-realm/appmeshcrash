package utils

import (
	"os"
	"strconv"

	"github.com/google/uuid"
)

// default configuration
const DefaultPort = "8090"
const DefaultText = "welcome to appmeshload"
const DefaultHost = "127.0.0.1"
const DefaultFport = "8090"
const DefaultFpath = "/get"
const DefaultDelay = "0"
const DefaultTimeout = "8"
const DefaultMultiHost = "http://127.0.0.1:8090/get"

func GetServerPort() string {
	port := os.Getenv("SERVER_PORT")
	if port != "" {
		return port
	}
	return DefaultPort
}

func GetForwardHost() string {
	host := os.Getenv("FORWARD_HOST")
	if host != "" {
		return host
	}
	return DefaultHost
}

func GetForwardPort() string {
	Fport := os.Getenv("FORWARD_PORT")
	if Fport != "" {
		return Fport
	}
	return DefaultFport
}

func GetForwardPath() string {
	Fpath := os.Getenv("FORWARD_PATH")
	if Fpath != "" {
		return Fpath
	}
	return DefaultFpath
}

func GetHostUID() string {
	return uuid.New().String()
}

func GetDelay() string {
	delay := os.Getenv("DELAY")
	if delay != "" {
		return delay
	}
	return DefaultDelay
}

func GetTimeout() string {
	timeout := os.Getenv("TIMEOUT")
	if timeout != "" {
		return timeout
	}
	return DefaultTimeout
}

func GetEcho() string {
	text := os.Getenv("TEXT")
	if text != "" {
		return text
	}
	return DefaultText
}

func GetMultiHost() string {
	multihost := os.Getenv("MULTIHOST")
	if multihost != "" {
		return multihost
	}
	return DefaultMultiHost
}

func GetLoadPrecision() int {
	loadprecision := os.Getenv("CPULOAD")
	precision, _ := strconv.Atoi(loadprecision)
	return precision * 2500000000
}
