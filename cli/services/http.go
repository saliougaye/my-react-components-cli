package services

import (
	"bytes"
	"encoding/json"
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

func (c HTTPClient) request(method, path, token string, body interface{}) (*http.Response, error) {

	bs, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, c.url+path, bytes.NewReader(bs))

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)

	res, err := c.client.Do(req)

	return res, err
}
