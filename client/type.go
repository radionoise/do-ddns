package client

import "fmt"

type apiError struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

func (e apiError) Error() string {
	return fmt.Sprintf("DigitalOcean API error. id: %v, message: %v", e.Id, e.Message)
}

type Domain struct {
	Name string `json:"name"`
}

func (d Domain) String() string {
	return fmt.Sprintf("{name: %v}", d.Name)
}

type DomainResponse struct {
	Domain Domain `json:"domain"`
}

type DomainListResponse struct {
	Domains []Domain `json:"domains"`
}

type DomainRecord struct {
	Id     int    `json:"id"`
	Type   string `json:"type"`
	Name   string `json:"name"`
	IpAddr string `json:"data"`
}

func (r DomainRecord) String() string {
	return fmt.Sprintf("{id: %v, type: %v, name: %v, ip: %v}", r.Id, r.Type, r.Name, r.IpAddr)
}

type DomainRecordResponse struct {
	DomainRecord DomainRecord `json:"domain_record"`
}

type DomainRecordListResponse struct {
	DomainRecords []DomainRecord `json:"domain_records"`
}
