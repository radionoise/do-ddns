package util

import (
	"errors"
	"fmt"
	"github.com/radionoise/do-ddns/client"
	"strings"
)

func GetSecondLevelName(domain string) (string, error) {
	tokens := strings.Split(domain, ".")

	if len(tokens) < 2 || len(tokens) > 3 {
		return "", errors.New(fmt.Sprintf("Invalid domain: %v", domain))
	}

	if len(tokens) == 2 {
		return domain, nil
	}

	return tokens[1] + "." + tokens[2], nil
}

func FindRecordByDomain(domain string, records []client.DomainRecord) (*client.DomainRecord, error) {
	tokens := strings.Split(domain, ".")

	for _, record := range records {
		if record.Type == "A" {
			if len(tokens) == 2 && record.Name == "@" {
				return &record, nil
			}

			if record.Name == tokens[0] {
				return &record, nil
			}
		}
	}

	return nil, errors.New(fmt.Sprintf("Cannot find record by domain: %v", domain))
}
