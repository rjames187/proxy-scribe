package handler

import (
	"io"
	"log"
	"net/http"
)

type ProxyHandler struct {
	ReqMethods map[string]func(string, string, io.ReadCloser) (*http.Response, error)
}

func (p *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host := "https://api.restful-api.dev"

	path := r.URL.Path
	contentType := r.Header.Get("Content-Type")
	reqBody := r.Body

	URL := host + path
	log.Printf("Requesting %s ...", URL)
	fn := p.ReqMethods[r.Method]
	resp, err := fn(URL, contentType, reqBody)
	if err != nil {
		log.Fatal(err)
	}

	returnBody, _ := io.ReadAll(resp.Body)
	w.WriteHeader(resp.StatusCode)
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.Write(returnBody)
}

func NewProxyHandler() *ProxyHandler {
	p := ProxyHandler{}
	return &p
}