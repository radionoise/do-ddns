package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	apiUrl     = "https://api.digitalocean.com/v2"
	domainsUrl = apiUrl + "/domains"
)

type Client struct {
	httpClient        *http.Client
	ipAddr            string
	hostname          string
	digitalOceanToken string
}

func New(ipAddr string, hostname string, digitalOceanToken string) *Client {
	return &Client{httpClient: &http.Client{}, ipAddr: ipAddr, hostname: hostname, digitalOceanToken: digitalOceanToken}
}

func (c Client) ListDomains() ([]Domain, error) {
	var domains Domains
	err := c.request(http.MethodGet, domainsUrl, nil, &domains)

	if err != nil {
		return nil, err
	}

	return domains.Domains, nil
}

func (c Client) setHeaders(request *http.Request) {
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %v", c.digitalOceanToken))
	request.Header.Add("Content-Type", "application/json")
}

func (c Client) request(method string, url string, body io.Reader, result interface{}) error {
	request, err := http.NewRequest(method, url, body)

	if err != nil {
		panic(err)
	}

	c.setHeaders(request)

	response, err := c.httpClient.Do(request)

	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		var apiError apiError
		err = json.NewDecoder(response.Body).Decode(&apiError)

		if err != nil {
			return err
		}

		return apiError
	}

	err = json.NewDecoder(response.Body).Decode(result)

	if err != nil {
		return err
	}

	return nil
}
