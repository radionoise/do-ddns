package main

import (
	"flag"
	"fmt"
	"github.com/radionoise/do-ddns/client"
	"github.com/radionoise/do-ddns/log"
	"github.com/radionoise/do-ddns/util"
	"os"
)

var (
	IpAddr            string
	Hostname          string
	DigitalOceanToken string
)

func main() {
	flag.StringVar(&IpAddr, "ip", "", "IP address")
	flag.StringVar(&Hostname, "host", "", "Hostname")
	flag.StringVar(&DigitalOceanToken, "token", "", "DigitalOcean access token")
	flag.Parse()

	if IpAddr == "" || Hostname == "" || DigitalOceanToken == "" {
		log.Error("Not enough parameters. See -h or --help for help")
		os.Exit(1)
	}

	doClient := client.New(IpAddr, Hostname, DigitalOceanToken)

	log.Debug("Getting domains")
	domains, err := doClient.ListDomains()
	errPanic(err)

	log.Debug(fmt.Sprintf("Found domains: %v", domains))
	parsedDomain, err := util.ParseDomain(Hostname)
	errPanic(err)

	createDomainIfNotExists(parsedDomain.Name, domains, doClient)

	log.Debug("Getting domain records")
	records, err := doClient.ListDomainRecords(parsedDomain.Name)
	errPanic(err)

	createOrUpdateRecord(IpAddr, *parsedDomain, records, doClient)
}

func errPanic(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func createDomainIfNotExists(secondLevel string, domains []client.Domain, doClient *client.Client) {
	found := false

	for _, val := range domains {
		if secondLevel == val.Name {
			log.Debug(fmt.Sprintf("Found existing domain: %v", val))
			found = true

			return
		}
	}

	if !found {
		log.Debug(fmt.Sprintf("Domain not found. Creating new domain: %v", secondLevel))
		response, err := doClient.CreateDomain(secondLevel)

		if err != nil {
			log.Panic(err)
		}

		log.Notice(fmt.Sprintf("Successfully created new domain: %v", response.Name))
	}
}

func createOrUpdateRecord(ipAddrd string, domain util.Domain, records []client.DomainRecord, doClient *client.Client) {
	var record *client.DomainRecord

	for _, val := range records {
		if val.Type == "A" && val.Name == domain.Record {
			record = &val

			break
		}
	}

	if record == nil {
		newRecord := client.DomainRecord{Type: "A", Name: domain.Record, IpAddr: ipAddrd}
		log.Debug(fmt.Sprintf("Domain record not found. Creating new record: %v", newRecord))
		record, err := doClient.CreateDomainRecord(domain.Name, newRecord)
		errPanic(err)
		log.Notice(fmt.Sprintf("Successfully created new domain %v with record: %v", domain.Name, record))

		return
	}

	record.IpAddr = ipAddrd

	record, err := doClient.UpdateDomainRecord(domain.Name, *record)
	errPanic(err)

	log.Notice(fmt.Sprintf("Successfully updated domain %v with record: %v", domain.Name, record))
}
