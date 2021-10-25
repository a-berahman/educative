package rest

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var client http.Client

//Initialize prepares basic items for REST process
func Initialize() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client = http.Client{
		Timeout:   5 * time.Second,
		Transport: tr,
	}
}

//PostApplicationJSON send request as a POST with JSON content type
func PostApplicationJSON(url string, body []byte) ([]byte, error) {
	request, err := makeRequest(url, body)
	if err != nil {
		return nil, err
	}

	response, err := doRequestAndGetData(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

//PostHTTPApplicationJSON send request as a POST with JSON content type by HTTP method
func PostHTTPApplicationJSON(url string, body []byte) ([]byte, error) {

	response, err := postHttpRequestByJSON(url, body)
	if err != nil {
		return nil, err
	}

	return response, nil
}
func makeRequest(url string, body []byte) (*http.Request, error) {
	request, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	return request, nil
}
func makeURLEncodingRequest(url string, data string) (*http.Request, error) {
	request, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return nil, err
	}
	return request, nil
}

// GetJSON send request as a GET with JSON content type
func GetJSON(url string, headers map[string]string) ([]byte, error) {
	request, err := makeGetRequest(url)
	if err != nil {
		return nil, err
	}

	headers["Content-Type"] = "application/json"
	addHeadersToRequest(request, headers)

	response, err := doRequestAndGetData(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func makeGetRequest(url string) (*http.Request, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func addHeadersToRequest(request *http.Request, headers map[string]string) {
	for key, value := range headers {
		request.Header.Add(key, value)
	}
}

func doRequestAndGetData(request *http.Request) ([]byte, error) {
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	bs, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

func postHttpRequestByJSON(url string, body []byte) ([]byte, error) {
	response, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	bs, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return bs, nil
}
