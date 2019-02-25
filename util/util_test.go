package util

import (
	"github.com/radionoise/do-ddns/test"
	"testing"
)

func TestGetSecondLevelName(t *testing.T) {
	domain, err := ParseDomain("example.com")
	test.AssertNil(t, err)
	test.AssertEquals(t, &Domain{Name: "example.com", Record: "@"}, domain)

	domain, err = ParseDomain("test.example.com")
	test.AssertNil(t, err)
	test.AssertEquals(t, &Domain{Name: "example.com", Record: "test"}, domain)

	domain, err = ParseDomain("com")
	test.AssertNotNil(t, err)
	test.AssertEquals(t, "Invalid domain: com", err.Error())

	domain, err = ParseDomain("a.b.example.com")
	test.AssertNotNil(t, err)
	test.AssertEquals(t, "Invalid domain: a.b.example.com", err.Error())
}
