package util

import (
	"errors"
	"fmt"
	"strings"
)

type Domain struct {
	Name   string
	Record string
}

func ParseDomain(domain string) (*Domain, error) {
	tokens := strings.Split(domain, ".")

	if len(tokens) < 2 || len(tokens) > 3 {
		return nil, errors.New(fmt.Sprintf("Invalid domain: %v", domain))
	}

	if len(tokens) == 2 {
		return &Domain{Name: domain, Record: "@"}, nil
	}

	return &Domain{Name: tokens[1] + "." + tokens[2], Record: tokens[0]}, nil
}
