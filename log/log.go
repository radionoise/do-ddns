package log

import (
	"fmt"
	"log/syslog"
)

const (
	SyslogTag = "do-ddns-client"
)

var writer *syslog.Writer

type Test struct {
	T string
	t int
}

func init() {
	syslogWriter, err := syslog.New(syslog.LOG_NOTICE, SyslogTag)

	if err != nil {
		panic(err)
	}

	writer = syslogWriter
}

func Notice(m string) {
	fmt.Println(m)
	err := writer.Notice(m)

	if err != nil {
		panic(err)
	}
}

func Error(m string) {
	fmt.Println(m)
	err := writer.Err(m)

	if err != nil {
		panic(err)
	}
}
