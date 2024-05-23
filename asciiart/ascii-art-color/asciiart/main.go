package main

import (
	"fmt"
	"os"
	"strings"

	asciiart "asciiart/functionFiles"
)

func main() {
	if len(os.Args) == 1 || len(os.Args) > 5 {
		fmt.Println("Usage: go run . [OPTION] [STRING] \n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
		return
	}

	var inputString, lettersToColor, color string

	switch len(os.Args) {
	case 2:
		inputString = os.Args[1]
		lettersToColor = inputString
		color = "resetCode"
	case 3:
		colorFlag := os.Args[1]
		inputString = os.Args[2]
		color = extractColor(colorFlag)
		if color == "" {
			printUsage()
			return
		}
		lettersToColor = inputString
	case 4:
		colorFlag := os.Args[1]
		lettersToColor = os.Args[2]
		inputString = os.Args[3]
		color = extractColor(colorFlag)
		if color == "" {
			printUsage()
			return
		}
	default:
		printUsage()
		return
	}

	characterMap, err := asciiart.CreateMap("../bannerfiles/standard.txt")
	if err != nil {
		fmt.Println("Error reading map:", err)
		return
	}

	asciiart.DisplayAsciiArtWithPartialColor(characterMap, inputString, lettersToColor, color)
}

func extractColor(colorFlag string) string {
	colorSplit := strings.Split(colorFlag, "=")
	if len(colorSplit) != 2 || colorSplit[0] != "--color" {
		return ""
	}
	return colorSplit[1]
}

func printUsage() {
	fmt.Println("Usage: go run . [OPTION] [STRING] \n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
}
