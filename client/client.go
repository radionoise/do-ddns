package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	apiUrl           = "https://api.digitalocean.com/v2"
	domainsUrl       = apiUrl + "/domains"
	domainRecordsUrl = domainsUrl + "/%v/records"
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
	var domains DomainListResponse
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

func (c Client) request(method string, url string, body interface{}, result interface{}) error {
	var jsonBody []byte
	var err error

	if body != nil {
		jsonBody, err = json.Marshal(body)

		if err != nil {
			return err
		}
	}

	var jsonBuffer *bytes.Buffer

	if body != nil {
		jsonBuffer = bytes.NewBuffer(jsonBody)
	}

	request, err := http.NewRequest(method, url, jsonBuffer)

	if err != nil {
		panic(err)
	}

	c.setHeaders(request)

	response, err := c.httpClient.Do(request)

	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
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

func (c Client) ListDomainRecords(domain string) ([]DomainRecord, error) {
	var records DomainRecordListResponse
	err := c.request(http.MethodGet, fmt.Sprintf(domainRecordsUrl, domain), nil, &records)

	if err != nil {
		return nil, err
	}

	return records.DomainRecords, nil
}

func (c Client) CreateDomain(domain string) (*Domain, error) {
	var newDomain DomainResponse
	err := c.request(http.MethodPost, domainsUrl, Domain{Name: domain}, &newDomain)

	if err != nil {
		return nil, err
	}

	return &newDomain.Domain, nil
}

func (c Client) CreateDomainRecord(domain string, record DomainRecord) (*DomainRecord, error) {
	var newRecord DomainRecordResponse
	err := c.request(http.MethodPost, fmt.Sprintf(domainRecordsUrl, domain), record, &newRecord)

	if err != nil {
		return nil, err
	}

	return &newRecord.DomainRecord, nil
}

func (c Client) UpdateDomainRecord(domain string, record DomainRecord) (*DomainRecord, error) {
	var newRecord DomainRecordResponse
	err := c.request(http.MethodPut, fmt.Sprintf(domainRecordsUrl, domain), record, &newRecord)

	if err != nil {
		return nil, err
	}

	return &newRecord.DomainRecord, nil
}
