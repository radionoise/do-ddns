package main

import (
	"flag"
	"fmt"
	"github.com/radionoise/do-ddns/client"
	"io/ioutil"
	"time"
)

var (
	IpAddr            string
	Hostname          string
	DigitalOceanToken string
	TzFile            string
)

func main() {
	flag.StringVar(&IpAddr, "ip", "", "IP address")
	flag.StringVar(&Hostname, "host", "", "Hostname")
	flag.StringVar(&DigitalOceanToken, "token", "", "DigitalOcean access token")
	flag.StringVar(&TzFile, "tz", "", "tzinfo file to override system timezone")
	flag.Parse()

	if IpAddr == "" || Hostname == "" || DigitalOceanToken == "" {
		panic("Not enough parameters. See -h or --help for help")
	}

	if TzFile != "" {
		overrideTimezone(TzFile)
	}

	doClient := client.New(IpAddr, Hostname, DigitalOceanToken)

	newDomain, err := doClient.CreateDomain("dan1.com")

	if err != nil {
		panic(err)
	}

	fmt.Println(newDomain.Name)
}

func overrideTimezone(tzFileName string) {
	fmt.Printf("Using timezone file: %v\n", tzFileName)

	result, err := ioutil.ReadFile(tzFileName)

	if err != nil {
		panic(fmt.Sprintf("Cannot open timezone file: %v", tzFileName))
	}

	location, err := time.LoadLocationFromTZData("", result)

	if err != nil {
		panic(fmt.Sprintf("Error loading timezone from tzif file: %v", err))
	}

	time.Local = location
}
