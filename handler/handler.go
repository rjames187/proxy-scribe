package handler

import (
	"io"
	"log"
	"net/http"
)

type ProxyHandler struct {
}

func (p *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host := "https://www.boredapi.com"
	path := r.URL.Path
	URL := host + path
	log.Printf("Request %s ...", URL)
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	w.WriteHeader(resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)
	log.Print(string(body))
	w.Write(body)
}