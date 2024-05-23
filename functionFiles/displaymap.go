package asciiart

import (
	"fmt"
	"strings"

	asciiart "asciiart/colorconversions"
)

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
		colorCode = asciiart.ConvertHexToAnsi(color)
		if colorCode == "" {
			fmt.Println("Error: Invalid hex color specified.")
			return
		}
	} else if strings.HasPrefix(color, "rgb(") {
		r, g, b, err := asciiart.ExtractRGBValues(color)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		colorCode = asciiart.ConvertRGBToAnsi(r, g, b)
	} else if strings.HasPrefix(color, "hsl(") {
		h, s, l, err := asciiart.ExtractHSLValues(color)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		r, g, b := asciiart.ConvertHSLToRGB(h, s, l)
		colorCode = asciiart.ConvertRGBToAnsi(r, g, b)
	} else {
		colorCode, ok = colorMap[color]
		if !ok {
			fmt.Println("Error: Invalid color specified.")
			return
		}
	}

	input = strings.ReplaceAll(input, "\\n", "\n")
	inputSlice := strings.Split(input, "\n")

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
