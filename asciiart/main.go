package main

import (
	"fmt"
	"os"
	"strings"

	asciiart "asciiart/functionFiles"
)

func main() {
	if len(os.Args) < 1 || len(os.Args) > 4 {
		fmt.Println("Usage: go run . --color=<color> <letters to be colored> \"something\"")
		return
	}
	var inputString string
	var lettersToColor string
	// var colorFlag string
	var color string
	if len(os.Args) == 2 {
		inputString = os.Args[1]
		lettersToColor = inputString
		color = "resetCode"
	}
	if len(os.Args) == 3 {
		colorFlag := os.Args[1]
		inputString = os.Args[2]
		colorSplit := strings.Split(colorFlag, "=")
		color = colorSplit[1]
		lettersToColor = inputString
		if len(colorSplit) != 2 || colorSplit[0] != "--color" {
			fmt.Println("Usage: go run . --color=<color> <letters to be colored> \"something\"")
			return
		}
	}
	if len(os.Args) == 4 {
		colorFlag := os.Args[1]
		inputString = os.Args[3]
		lettersToColor = os.Args[2]
		colorSplit := strings.Split(colorFlag, "=")
		color = colorSplit[1]
		if len(colorSplit) != 2 || colorSplit[0] != "--color" {
			fmt.Println("Usage: go run . --color=<color> <letters to be colored> \"something\"")
			return
		}

	}

	// Extract the letters to be colored and the rest of the string
	// lettersToColor, inputString = extractLettersToColor(args)

	characterMap, err := asciiart.CreateMap("standard.txt")
	if err != nil {
		fmt.Println("Error reading map:", err)
		return
	}

	asciiart.DisplayAsciiArtWithPartialColor(characterMap, inputString, lettersToColor, color)
}
