package main

import (
	"log"
	"net/http"

	"go-rest-server-template/config"
	"go-rest-server-template/server"
)

var configObj *config.ConfigObject

func init() {
	configObj = config.GetInstance()
}

func main() {
	router := server.NewServerRouter()
	log.Fatal(http.ListenAndServe("0.0.0.0:"+configObj.Port, router))
}
