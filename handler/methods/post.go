package methods

import (
	"io"
	"log"
	"net/http"
)

func HandlePost(w http.ResponseWriter, r *http.Request) {
	host := "https://api.restful-api.dev"
	path := r.URL.Path
	URL := host + path
	contentType := r.Header.Get("Content-Type")
	reqBody := r.Body
	log.Printf("Request %s ...", URL)
	resp, err := http.Post(URL, contentType, reqBody)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	w.WriteHeader(resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)
	log.Print(string(body))
	w.Write(body)
}