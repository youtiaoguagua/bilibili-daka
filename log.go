package main

import (
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	writer1 := os.Stdout
	var writer2 *os.File
	logFile := "log.txt"
	if _, err := os.Stat(logFile); err == nil {
		writer2, _ = os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE, 0755)
	} else {
		writer2, _ = os.Create(logFile)
	}
	log.SetOutput(io.MultiWriter(writer1, writer2))

}
