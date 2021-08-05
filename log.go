package main

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

func init() {
	log.SetFormatter(&Log{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	writer1 := os.Stdout
	var writer2 *os.File
	logFile := "./log.txt"
	writer2, _ = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE, 0777)
	log.SetOutput(io.MultiWriter(writer1, writer2))
}

type Log struct {
}

func (l Log) Format(entry *log.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	newLog := fmt.Sprintf("%s\n", entry.Message)

	b.WriteString(newLog)
	return b.Bytes(), nil
}
