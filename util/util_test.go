package util

import (
	"github.com/radionoise/do-ddns/client"
	"github.com/radionoise/do-ddns/test"
	"testing"
)

func TestGetSecondLevelName(t *testing.T) {
	domain, err := GetSecondLevelName("example.com")
	test.AssertNil(t, err)
	test.AssertEquals(t, "example.com", domain)

	domain, err = GetSecondLevelName("test.example.com")
	test.AssertNil(t, err)
	test.AssertEquals(t, "example.com", domain)

	domain, err = GetSecondLevelName("com")
	test.AssertNotNil(t, err)
	test.AssertEquals(t, "Invalid domain: com", err.Error())

	domain, err = GetSecondLevelName("a.b.example.com")
	test.AssertNotNil(t, err)
	test.AssertEquals(t, "Invalid domain: a.b.example.com", err.Error())
}

func TestFindRecordByDomain(t *testing.T) {
	records := []client.DomainRecord{
		{Id: 1, Type: "NS", Name: "@"},
		{Id: 2, Type: "A", Name: "@"},
		{Id: 3, Type: "A", Name: "subdomain"},
		{Id: 4, Type: "MX", Name: "@"},
	}

	foundRecord, err := FindRecordByDomain("example.com", records)
	test.AssertNil(t, err)
	test.AssertEquals(t, &client.DomainRecord{Id: 2, Type: "A", Name: "@"}, foundRecord)

	foundRecord, err = FindRecordByDomain("subdomain.example.com", records)
	test.AssertNil(t, err)
	test.AssertEquals(t, &client.DomainRecord{Id: 3, Type: "A", Name: "subdomain"}, foundRecord)

	foundRecord, err = FindRecordByDomain("notexists.example.com", records)
	test.AssertNotNil(t, foundRecord)
	test.AssertEquals(t, "Cannot find record by domain: notexists.example.com", err.Error())

	foundRecord, err = FindRecordByDomain("example", records)
	test.AssertNotNil(t, foundRecord)
	test.AssertEquals(t, "Cannot find record by domain: example", err.Error())
}
