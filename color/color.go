package color

import "math"

type Hsv struct {
	Hue float64
	Saturation float64
	Value float64
}

func Color(red uint32, green uint32, blue uint32) uint32 {
	var rgb uint32

	rgb = green
	rgb = (rgb << 8) + red
	rgb = (rgb << 8) + blue

	return rgb
}

func HsvToColor(hue float64, saturation float64, value float64) uint32 {
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