package main

import (
	"log/syslog"
	"log"
	"ara/color"
	"flag"
)

var config Config = getConfig()

func main() {
	logger, err := syslog.New(syslog.LOG_ERR, "ara")

	if err != nil {
		log.Fatal("Error writing syslog!")
	}

	hue := flag.Float64("hue", float64(0), "Hue channel")
	saturation := flag.Float64("saturation", float64(1), "Saturation channel")
	value := flag.Float64("value", float64(1), "Value channel")

	flag.Parse()

	hsv := color.Hsv{*hue, *saturation, *value}

	fb := make(chan string)

	logger.Info("Starting Ara...")
	go driver(logger, fb, &hsv)

	logger.Info(<-fb)
	logger.Info("Ara is up and running!")

	initShell(&hsv)
}