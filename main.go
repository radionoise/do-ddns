package main

import (
	"fmt"
	"github.com/radionoise/do-ddns/log"
	"io/ioutil"
	"time"
)

var ()

func main() {

	fmt.Println(time.Now())
	log.Notice("Hello world!")
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
