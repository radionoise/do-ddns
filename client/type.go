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

type Domains struct {
	Domains []Domain `json:"domains"`
}
