package asciiart

import (
	"fmt"
	"strconv"
	"strings"

	asciiart "asciiart/colorconversions"
)

// ExtractRGBValues extracts RGB values from a string in the format "rgb(r, g, b)".
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
	if r < 0 || r > 255 || g < 0 || g > 255 || b < 0 || b > 255 {
		return 0, 0, 0, fmt.Errorf("rgb values must be between 0 and 255")
	}
	return r, g, b, nil
}

// ExtractHSLValues extracts HSL values from a string in the format "hsl(h, s, l)".
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
	h, err1 := strconv.ParseFloat(strings.TrimSpace(values[0]), 64)
	s, err2 := strconv.ParseFloat(strings.TrimSuffix(strings.TrimSpace(values[1]), "%"), 64)
	l, err3 := strconv.ParseFloat(strings.TrimSuffix(strings.TrimSpace(values[2]), "%"), 64)
	if err1 != nil || err2 != nil || err3 != nil {
		return 0, 0, 0, fmt.Errorf("invalid hsl values")
	}
	if h < 0 || h > 360 || s < 0 || s > 100 || l < 0 || l > 100 {
		return 0, 0, 0, fmt.Errorf("hsl values must be within valid ranges")
	}
	return h, s, l, nil
}

// DisplayAsciiArtWithPartialColor displays ASCII art with specified letters in color.
func DisplayAsciiArtWithPartialColor(characterMap map[rune][]string, input, lettersToColor, color string) {
	if len(characterMap) == 0 {
		fmt.Println("Error: Character map is empty. Please provide a valid character map.")
		return
	}

	resetCode := "\033[0m"

	colorCode := determineColorCode(color)
	if colorCode == "" {
		fmt.Println("Error: Invalid color specified.")
		return
	}

	input = strings.ReplaceAll(input, "\\n", "\n")
	lines := strings.Split(input, "\n")

	for _, line := range lines {
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

// determineColorCode returns the appropriate ANSI color code based on the input color.
func determineColorCode(color string) string {
	colorMap := map[string]string{
		"cyan":          "\033[36m",
		"grey":          "\033[37m",
		"yellow":        "\033[33m",
		"black":         "\x1b[30m",
		"red":           "\x1b[31m",
		"green":         "\x1b[32m",
		"blue":          "\x1b[34m",
		"magenta":       "\x1b[35m",
		"white":         "\x1b[37m",
		"pink":          "\x1b[38;5;205m",
		"brown":         "\x1b[38;5;94m",
		"orange":        "\x1b[38;5;214m",
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

	if strings.HasPrefix(color, "#") {
		return asciiart.ConvertHexToAnsi(color)
	} else if strings.HasPrefix(color, "rgb(") {
		r, g, b, err := ExtractRGBValues(color)
		if err != nil {
			return ""
		}
		return asciiart.ConvertRGBToAnsi(r, g, b)
	} else if strings.HasPrefix(color, "hsl(") {
		h, s, l, err := ExtractHSLValues(color)
		if err != nil {
			return ""
		}
		r, g, b := asciiart.ConvertHSLToRGB(h, s, l)
		return asciiart.ConvertRGBToAnsi(r, g, b)
	}

	return colorMap[color]
}
