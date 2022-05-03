package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type HTTPClient struct {
	client *http.Client
	url    string
}

func NewHTTPClient(url string) HTTPClient {
	return HTTPClient{
		client: &http.Client{},
		url:    url,
	}
}

func (c HTTPClient) request(method, path string, headers map[string]string, body interface{}) (*http.Response, error) {

	bs, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, c.url+path, bytes.NewBuffer(bs))

	fmt.Println()

	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	res, err := c.client.Do(req)

	return res, err
}
