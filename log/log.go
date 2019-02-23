package log

import (
	"fmt"
	"log/syslog"
	"os"
)

const (
	SyslogTag = "do-ddns"
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

func Debug(m string) {
	fmt.Println(m)
	err := writer.Debug(m)

	if err != nil {
		panic(err)
	}
}

func Notice(m string) {
	fmt.Println(m)
	err := writer.Notice(m)

	if err != nil {
		panic(err)
	}
}

func Error(m string) {
	_, err := fmt.Fprint(os.Stderr, m)

	if err != nil {
		panic(err)
	}

	err = writer.Err(m)

	if err != nil {
		panic(err)
	}
}
