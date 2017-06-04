package main

import (
	"log/syslog"
	"rpi_ws281x/golang/ws2811"
	"fmt"
	"os"
	"time"
	"ara/color"
)

func driver(logger *syslog.Writer, fb chan string, hsv *color.Hsv) {
	if err := ws2811.Init(18, 300, 255); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fb <- "Initialized driver"

	for b := 0; b >= 0; b++ {
		for i := 0; i < config.Length; i++ {
			color := color.HsvToColor(hsv.Hue, hsv.Saturation, hsv.Value)

			ws2811.SetLed(i, color)
		}

		ws2811.Render()

		time.Sleep(time.Second / time.Duration(config.Fps))
	}
}