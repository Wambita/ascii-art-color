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
	//color map ansi
	colorMap := asciiart.AnsiColorMap()

	var colorCode string
	var ok bool

	//color conversions rgb, hex, hsl
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
