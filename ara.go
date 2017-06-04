package main

import (
	"log/syslog"
	"log"
	"ara/color"
)

var config Config = getConfig()

var hsv color.Hsv = color.Hsv{0.0, 1.0, 1.0}

func main() {
	logger, err := syslog.New(syslog.LOG_ERR, "ara")

	if err != nil {
		log.Fatal("Error writing syslog!")
	}

	fb := make(chan string)

	logger.Info("Starting Ara...")
	go driver(logger, fb)

	logger.Info(<-fb)
	logger.Info("Ara is up and running!")

	initShell()
}