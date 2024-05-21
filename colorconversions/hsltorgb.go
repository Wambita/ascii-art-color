package asciiart

import "math"

// ConvertHSLToRGB converts HSL values to RGB values.
func ConvertHSLToRGB(h, s, l float64) (int, int, int) {
	// normalize saturation and hue values
	s = s / 100
	l = l / 100

	// Chroma (c): Represents the difference between the maximum and minimum RGB components.
	// X (x): Represents an intermediate value used to calculate the RGB components. It depends on the hue and chroma
	// M (m): Represents the amount to be added to each RGB component to match the lightness:
	c := (1 - abs(2*l-1)) * s
	x := c * (1 - abs(math.Mod(h/60, 2)-1))
	m := l - c/2

	var r, g, b float64

	// assign values based on hue
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

	// adjust lightness and rgb values
	r = (r + m) * 255
	g = (g + m) * 255
	b = (b + m) * 255

	return int(r), int(g), int(b)
}

// abs is a helper function to compute absolute value.
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
