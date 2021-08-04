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
	writer2, _ := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0755)
	log.SetOutput(io.MultiWriter(writer1, writer2))

}
