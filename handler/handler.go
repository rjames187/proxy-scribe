package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"proxy-scribe/spec"
)

type ProxyHandler struct {
	Doc *spec.Spec
}

func (p *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host := "https://api.restful-api.dev"

	path := r.URL.Path
	queryParams := r.URL.Query()
	contentType := r.Header.Get("Content-Type")
	method := r.Method
	reqBody := r.Body
	defer reqBody.Close()
	var buffer bytes.Buffer
	var reqBodyData map[string]interface{}

	io.Copy(&buffer, reqBody)
	decoder := json.NewDecoder(bytes.NewReader(buffer.Bytes()))
	if err := decoder.Decode(&reqBodyData); err != nil {
		fmt.Print("Error! Problem decoding JSON body")
		os.Exit(1)
	}

	// record request information
	var doc *spec.Spec
	if p.Doc == nil {
		doc = spec.NewSpec()
		p.Doc = doc
	} else {
		doc = p.Doc
	}
	doc.ReadInPath(path)
	doc.ReadInMethod(path, method)
	doc.ReadInReq(path, method, reqBodyData)
	doc.ReadInQParams(path, method, queryParams)

	URL := host + path
	log.Printf("\n\tRequesting %s ...", URL)
	fmt.Print("ProxyScribe > ")

	reader := bytes.NewReader(buffer.Bytes())
	resp, err := sendReq(method, URL, contentType, reader)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var resBodyData map[string]interface{}
	buffer = bytes.Buffer{}
	io.Copy(&buffer, resp.Body)
	decoder = json.NewDecoder(bytes.NewReader(buffer.Bytes()))
	if err := decoder.Decode(&resBodyData); err != nil {
		fmt.Print("Error! Problem decoding JSON body")
		os.Exit(1)
	}
	doc.ReadInRes(path, method, resp.StatusCode, resBodyData)

	w.WriteHeader(resp.StatusCode)
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.Write(buffer.Bytes())
}