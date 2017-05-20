package main

import (
	"fmt"
	"math"
	"os"
	"rpi_ws281x/golang/ws2811"
	"time"
)

var config Config = getConfig()

func main() {
	if err := ws2811.Init(18, 300, 255); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for b := 0; b >= 0; b++ {
		for i := 0; i < config.Length; i++ {
			color := Hsv(float64(i*360/config.Length)+float64(b), 1.0, 1.0)

			ws2811.SetLed(i, color)
		}

		ws2811.Render()

		time.Sleep(time.Second / time.Duration(config.Fps))
	}
}

func Color(red uint32, green uint32, blue uint32) uint32 {
	var rgb uint32

	rgb = green
	rgb = (rgb << 8) + red
	rgb = (rgb << 8) + blue

	return rgb
}

func Hsv(hue float64, saturation float64, value float64) uint32 {
	var red, green, blue float64

	hue = math.Mod(hue, 360) / 360

	i := math.Floor(hue * 6)
	f := hue*6.0 - i
	p := value * (1 - saturation)
	q := value * (1 - f*saturation)
	t := value * (1 - (1-f)*saturation)

	switch int(i) % 6 {
	case 0:
		red = value
		green = t
		blue = p
	case 1:
		red = q
		green = value
		blue = p
	case 2:
		red = p
		green = value
		blue = t
	case 3:
		red = p
		green = q
		blue = value
	case 4:
		red = t
		green = p
		blue = value
	case 5:
		red = value
		green = p
		blue = q
	}

	return Color(uint32(red*255), uint32(green*255), uint32(blue*255))
}
