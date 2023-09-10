package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"proxy-scribe/spec"
)

type ProxyHandler struct {
}

func (p *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host := "https://api.restful-api.dev"

	path := r.URL.Path
	contentType := r.Header.Get("Content-Type")
	method := r.Method
	reqBody := r.Body
	defer reqBody.Close()
	var reqBodyData map[string]interface{}

	decoder := json.NewDecoder(reqBody)
	if err := decoder.Decode(&reqBodyData); err != nil {
		fmt.Print("Error! Problem decoding JSON body")
		os.Exit(1)
	}

	// record request information
	doc := spec.NewSpec()
	doc.ReadInPath(path)
	doc.ReadInMethod(path, method)
	doc.ReadInReq(path, method, reqBodyData)


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