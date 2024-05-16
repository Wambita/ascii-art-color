package asciiart

import (
	"fmt"
	"strings"
)

func DisplayAsciiArt(characterMap map[rune][]string, input, color string) {
	// Check if the character map is empty
	if len(characterMap) == 0 {
		fmt.Println("Error: Character map is empty. Please provide a valid character map.")
		return
	}
	// Split the input string by "\n" to handle newline characters.
	inputSlice := strings.Split(input, "\\n")
	count := 0
	colorMap := map[string]string{
		"cyan":   "\033[36m",
		"grey":   "\033[37m",
		"yellow": "\033[33m",
		"black" : "\x1b[30m",
		"red": "\x1b[31m",
		"green" :"\x1b[32m",
		"blue":"\x1b[34m",
		"magneta":"\x1b[35m",
		"white" : "\x1b[37m",

	}

	// Iterate through the split input segments.
	// If the segment is empty (indicating a newline), it prints a newline character, ensuring proper formatting.
	/* If the segment contains characters:
	   Iterates through each character in the segment.
	   Retrieves the corresponding ASCII art lines for the character from the character map.
	   Prints each line of ASCII art for the character.
	   If the character is not found in the map, it prints an error message and returns.*/

	for _, value := range inputSlice {
		// var Cyan = "\033[36m"

		if value == "" {
			count++
			if count < len(inputSlice) {
				fmt.Println()
			}
		} else {
			color, str := colorMap[color]
			for i := 0; i < 8; i++ {
				for _, char := range value {
					line, ok := characterMap[char]
					if ok && str {
						fmt.Print(color + line[i])
					} else {
						fmt.Println("Characters not found")
						return
					}
				}
				fmt.Println()
			}
		}
	}
}
