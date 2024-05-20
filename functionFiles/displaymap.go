package asciiart

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// ConvertHexToAnsi converts a hex color code to the nearest ANSI 256 color code.
func ConvertHexToAnsi(hex string) string {
	if len(hex) != 7 || hex[0] != '#' {
		return ""
	}

	r, _ := strconv.ParseInt(hex[1:3], 16, 64)
	g, _ := strconv.ParseInt(hex[3:5], 16, 64)
	b, _ := strconv.ParseInt(hex[5:7], 16, 64)

	return ConvertRGBToAnsi(int(r), int(g), int(b))
}

// ConvertRGBToAnsi converts RGB values to the nearest ANSI 256 color code.
func ConvertRGBToAnsi(r, g, b int) string {
	color := 16 + (36 * (r / 51)) + (6 * (g / 51)) + (b / 51)
	return fmt.Sprintf("\033[38;5;%dm", color)
}

// ConvertHSLToRGB converts HSL values to RGB values.
func ConvertHSLToRGB(h, s, l float64) (int, int, int) {
	s /= 100
	l /= 100
	c := (1 - abs(2*l-1)) * s
	x := c * (1 - abs(math.Mod(h/60, 2)-1))
	m := l - c/2

	var r, g, b float64

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

func ExtractRGBValues(rgb string) (int, int, int, error) {
	if !strings.HasPrefix(rgb, "rgb(") || !strings.HasSuffix(rgb, ")") {
		return 0, 0, 0, fmt.Errorf("invalid rgb format")
	}
	rgb = strings.TrimPrefix(rgb, "rgb(")
	rgb = strings.TrimSuffix(rgb, ")")
	values := strings.Split(rgb, ",")
	if len(values) != 3 {
		return 0, 0, 0, fmt.Errorf("invalid rgb format")
	}
	r, err1 := strconv.Atoi(strings.TrimSpace(values[0]))
	g, err2 := strconv.Atoi(strings.TrimSpace(values[1]))
	b, err3 := strconv.Atoi(strings.TrimSpace(values[2]))
	if err1 != nil || err2 != nil || err3 != nil {
		return 0, 0, 0, fmt.Errorf("invalid rgb values")
	}
	if r < 0 || r > 255 {
		return 0, 0, 0, fmt.Errorf("red value must be between 0 and 255")
	}
	if g < 0 || g > 255 {
		return 0, 0, 0, fmt.Errorf("green value must be between 0 and 255")
	}
	if b < 0 || b > 255 {
		return 0, 0, 0, fmt.Errorf("blue value must be between 0 and 255")
	}

	return r, g, b, nil
}

// ExtractHSLValues extracts the h, s, l values from an hsl string.
func ExtractHSLValues(hsl string) (float64, float64, float64, error) {
	if !strings.HasPrefix(hsl, "hsl(") || !strings.HasSuffix(hsl, ")") {
		return 0, 0, 0, fmt.Errorf("invalid hsl format")
	}
	hsl = strings.TrimPrefix(hsl, "hsl(")
	hsl = strings.TrimSuffix(hsl, ")")
	values := strings.Split(hsl, ",")
	if len(values) != 3 {
		return 0, 0, 0, fmt.Errorf("invalid hsl format")
	}
	h, err1 := strconv.ParseFloat(strings.TrimSuffix(strings.TrimSpace(values[0]), "%"), 64)
	s, err2 := strconv.ParseFloat(strings.TrimSuffix(strings.TrimSpace(values[1]), "%"), 64)
	l, err3 := strconv.ParseFloat(strings.TrimSuffix(strings.TrimSpace(values[2]), "%"), 64)
	if err1 != nil || err2 != nil || err3 != nil {
		return 0, 0, 0, fmt.Errorf("invalid hsl values")
	}
	if h < 0 || h > 360 {
		return 0, 0, 0, fmt.Errorf("hue value must be between 0 and 360")
	}
	if s < 0 || s > 100 {
		return 0, 0, 0, fmt.Errorf("saturation value must be between 0 and 100")
	}
	if l < 0 || l > 100 {
		return 0, 0, 0, fmt.Errorf("lightness value must be between 0 and 100")
	}
	return h, s, l, nil
}

func DisplayAsciiArtWithPartialColor(characterMap map[rune][]string, input, lettersToColor, color string) {
	if len(characterMap) == 0 {
		fmt.Println("Error: Character map is empty. Please provide a valid character map.")
		return
	}

	resetCode := "\033[0m"

	colorMap := map[string]string{
		"cyan":          "\033[36m",
		"grey":          "\033[37m",
		"yellow":        "\033[33m",
		"black":         "\x1b[30m",
		"red":           "\x1b[31m",
		"green":         "\x1b[32m",
		"blue":          "\x1b[34m",
		"magneta":       "\x1b[35m",
		"white":         "\x1b[37m",
		"pink":          "\x1b[38;5;205m", // ANSI code for pink
		"brown":         "\x1b[38;5;94m",  // ANSI code for brown
		"orange":        "\x1b[38;5;214m", // ANSI code for orange
		"lime":          "\x1b[38;5;10m",
		"brightBlack":   "\x1b[90m",
		"brightRed":     "\x1b[91m",
		"brightGreen":   "\x1b[92m",
		"brightYellow":  "\x1b[93m",
		"brightBlue":    "\x1b[94m",
		"brightMagenta": "\x1b[95m",
		"brightCyan":    "\x1b[96m",
		"brightWhite":   "\x1b[97m",
		"resetCode":     "\033[0m",
	}

	var colorCode string
	var ok bool

	if strings.HasPrefix(color, "#") {
		colorCode = ConvertHexToAnsi(color)
		if colorCode == "" {
			fmt.Println("Error: Invalid hex color specified.")
			return
		}
	} else if strings.HasPrefix(color, "rgb(") {
		r, g, b, err := ExtractRGBValues(color)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		colorCode = ConvertRGBToAnsi(r, g, b)
	} else if strings.HasPrefix(color, "hsl(") {
		h, s, l, err := ExtractHSLValues(color)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		r, g, b := ConvertHSLToRGB(h, s, l)
		colorCode = ConvertRGBToAnsi(r, g, b)
	} else {
		colorCode, ok = colorMap[color]
		if !ok {
			fmt.Println("Error: Invalid color specified.")
			return
		}
	}

	inputSlice := strings.Split(input, "\\n")

	for _, line := range inputSlice {
		if line == "" {
			fmt.Println()
		} else {
			for i := 0; i < 8; i++ {
				for _, char := range line {
					artLines, ok := characterMap[char]
					if ok {
						if strings.ContainsRune(lettersToColor, char) {
							fmt.Print(colorCode + artLines[i] + resetCode)
						} else {
							fmt.Print(artLines[i])
						}
					} else {
						fmt.Printf("Error: Character '%c' not found in map.\n", char)
						return
					}
				}
				fmt.Println()
			}
		}
	}
}
