package handler

import (
	"io"
	"net/http"
)

func sendReq(m string, URL string, contentType string, body io.ReadCloser) (*http.Response, error) {
	req, err := http.NewRequest(m, URL, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}