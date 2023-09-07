package main

import (
	"log"
	"net/http"
	"proxy-scribe/handler"
)

func main() {
	log.Print("Proxy is listening")
	controller := &handler.ProxyHandler{}
	log.Fatal(http.ListenAndServe(":4000", controller))
}