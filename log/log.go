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
	syslogWrite(m, writer.Debug)
}

func syslogWrite(m string, writeFunc func(string) error) {
	err := writeFunc(m)

	if err != nil {
		panic(err)
	}
}

func Notice(m string) {
	fmt.Println(m)
	syslogWrite(m, writer.Notice)
}

func Error(m string) {
	_, err := fmt.Fprint(os.Stderr, m)

	if err != nil {
		panic(err)
	}

	syslogWrite(m, writer.Err)
}

func Panic(err error) {
	syslogWrite(err.Error(), writer.Err)
	panic(err)
}
