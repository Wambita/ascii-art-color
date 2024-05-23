package asciiart

import "math"

// ConvertHSLToRGB converts HSL (Hue, Saturation, Lightness) values to RGB (Red, Green, Blue) values.
// The input values are:
// - h: Hue, ranges from 0 to 360 degrees.
// - s: Saturation, ranges from 0 to 100 percent.
// - l: Lightness, ranges from 0 to 100 percent.
// The function returns three integers representing the RGB values, each ranging from 0 to 255.
func ConvertHSLToRGB(h, s, l float64) (int, int, int) {
	// Normalize saturation and lightness values to be within 0 and 1.
	s /= 100
	l /= 100

	// Chroma: the maximum and minimum RGB components.
	c := (1 - math.Abs(2*l-1)) * s

	// X is an intermediate value used to calculate the RGB components.
	// It depends on the hue and chroma.
	x := c * (1 - math.Abs(math.Mod(h/60, 2)-1))

	// M  is the amount to be added to each RGB component to match the lightness.
	m := l - c/2

	// Initialize rgb values.
	var r, g, b float64

	// Determine RGB values based on the hue.
	switch {
	case h < 60:
		r, g, b = c, x, 0
	case h < 120:
		r, g, b = x, c, 0
	case h < 180:
		r, g, b = 0, c, x
	case h < 240:
		r, g, b = 0, x, c
	case h < 300:
		r, g, b = x, 0, c
	default:
		r, g, b = c, 0, x
	}

	// Adjust RGB values by adding m and converting to the range 0-255.
	return int((r + m) * 255), int((g + m) * 255), int((b + m) * 255)
}
