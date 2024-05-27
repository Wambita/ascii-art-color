package main

import (
	"fmt"
	"os"
	"strings"

	asciiart "asciiart/functionFiles"
)

func main() {
	// check length of arguments
	if len(os.Args) == 1 || len(os.Args) > 5 {
		fmt.Println("Usage: go run . [OPTION] [STRING] \n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
		return
	}
	var (
		inputString    string
		lettersToColor string
		color          string
	)
	// case: no colors selected
	if len(os.Args) == 2 {
		inputString = os.Args[1]
		lettersToColor = inputString
		color = "resetCode"
	}
	// case color selected but no specified letters to be colored
	if len(os.Args) == 3 {
		colorFlag := os.Args[1]
		inputString = os.Args[2]
		if !(strings.HasPrefix(colorFlag, "--color=")) {
			fmt.Println("Usage: go run . [OPTION] [STRING] \n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
			return
		}
		colorSplit := strings.Split(colorFlag, "=")
		color = colorSplit[1]
		lettersToColor = inputString
		if len(colorSplit) != 2 || colorSplit[0] != "--color" {
			fmt.Println("Usage: go run . [OPTION] [STRING] \n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
			return
		}
	}
	// case color selected , letters to be colored specified and input string
	if len(os.Args) == 4 {
		colorFlag := os.Args[1]
		inputString = os.Args[3]
		lettersToColor = os.Args[2]
		if !(strings.HasPrefix(colorFlag, "--color=")) {
			fmt.Println("Usage: go run . [OPTION] [STRING] \n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
			return
		}
		colorSplit := strings.Split(colorFlag, "=")
		color = colorSplit[1]
		if len(colorSplit) != 2 || colorSplit[0] != "--color" {
			fmt.Println("Usage: go run . [OPTION] [STRING] \n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
			return
		}
	}

	// Extract the letters to be colored and the rest of the string
	// lettersToColor, inputString = extractLettersToColor(args)

	characterMap, err := asciiart.CreateMap("../bannerfiles/standard.txt")
	if err != nil {
		fmt.Println("Error reading map:", err)
		return
	}

	asciiart.DisplayAsciiArtWithPartialColor(characterMap, inputString, lettersToColor, color)
}
