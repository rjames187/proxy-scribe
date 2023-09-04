package main

import (
	"log"
	"net/http"
	"proxy-scribe/handler"
)

func main() {
	log.Print("Proxy is listening")
	mm := handler.NewMethodMux()
	log.Fatal(http.ListenAndServe(":4000", mm))
}