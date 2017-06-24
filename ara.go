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
	logger.Info("Starting Ara...")

	if err != nil {
		log.Fatal("Error writing syslog!")
	}

	hue := flag.String("hue", "0", "Hue channel")
	saturation := flag.String("saturation", "1", "Saturation channel")
	value := flag.String("value", "1", "Value channel")

	flag.Parse()

	hsv := color.Hsv{}

	if hueFn, err := compile(*hue); err == nil {
		hsv.Hue = hueFn
	}

	if saturationFn, err := compile(*saturation); err == nil {
		hsv.Saturation = saturationFn
	}

	if valueFn, err := compile(*value); err == nil {
		hsv.Value = valueFn
	}

	go driver(logger, &hsv)
	go server(logger, &hsv)

	logger.Info("Ara is up and running!")

	initShell(logger, &hsv)
}
