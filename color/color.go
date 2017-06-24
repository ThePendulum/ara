package color

import (
	"github.com/yuin/gopher-lua"
	"math"
)

type Hsv struct {
	Hue *lua.LFunction
	Saturation *lua.LFunction
	Value *lua.LFunction
}

func HsvToRgb(hue float64, saturation float64, value float64) (float64, float64, float64) {
	var red, green, blue float64

	hue = math.Abs(math.Mod(hue, 360)) / 360
	saturation = math.Max(float64(0), math.Min(float64(1), saturation))
	value = math.Max(float64(0), math.Min(float64(1), value))

	i := math.Floor(hue * 6)
	f := hue * 6.0 - i
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

	return float64(red * 255), float64(green * 255), float64(blue * 255)
}

func HsvToColor(hue float64, saturation float64, value float64) uint32 {
	red, green, blue := HsvToRgb(hue, saturation, value)

	return Color(uint32(red), uint32(green), uint32(blue))
}

func Color(red uint32, green uint32, blue uint32) uint32 {
	var rgb uint32

	rgb = green
	rgb = (rgb << 8) + red
	rgb = (rgb << 8) + blue

	return rgb
}
