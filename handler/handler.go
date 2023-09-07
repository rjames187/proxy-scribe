package handler

import (
	"io"
	"log"
	"net/http"
)

type ProxyHandler struct {
}

func (p *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host := "https://api.restful-api.dev"

	path := r.URL.Path
	contentType := r.Header.Get("Content-Type")
	reqBody := r.Body
	method := r.Method

	URL := host + path
	log.Printf("Requesting %s ...", URL)
	resp, err := sendReq(method, URL, contentType, reqBody)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	returnBody, _ := io.ReadAll(resp.Body)
	w.WriteHeader(resp.StatusCode)
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.Write(returnBody)
}