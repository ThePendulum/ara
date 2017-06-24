package main

import (
	"log/syslog"
	"rpi_ws281x/golang/ws2811"
	"time"
	"ara/color"
	"os"
	"fmt"
)

func driver(logger *syslog.Writer, hsv *color.Hsv) {
	context := map[string]int{
		"length": config.Length,
	}

	if err := ws2811.Init(18, 300, 255); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for b := 0; b >= 0; b++ {
		context["beat"] = b

		for i := 0; i < config.Length; i++ {
			context["index"] = i

			result := color.HsvToColor(call(hsv.Hue, context), call(hsv.Saturation, context), call(hsv.Value, context))

			ws2811.SetLed(i, result)
		}

		ws2811.Render()

		time.Sleep(time.Second / time.Duration(config.Fps))
	}
}
