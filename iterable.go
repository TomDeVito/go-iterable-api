package iterable

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	ROOT_URI = "https://api.iterable.com:443/api/"
)

type AuthOptions struct {
	API_KEY string // Iterable API key from iterable settings api key
}

type Client struct {
	Options *AuthOptions
	Client  *http.Client
}

type IterableAPI interface {
	RegisterDeviceToken(registerDeviceTokenRequest *RegisterDeviceTokenRequest) error
}

func NewClient(options *AuthOptions, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		Options: options,
		Client:  httpClient,
	}
}

func (client *Client) iterableCall(requestType string, url string, requestBody []byte) ([]byte, error) {

	fullUrl := fmt.Sprintf("%s%s?api_key=%s", ROOT_URI, url, client.Options.API_KEY)
	req, err := http.NewRequest(requestType, fullUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	response, err := client.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if response.StatusCode > 300 {
		return nil, fmt.Errorf("Failed API call: %s", response.Status)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}
