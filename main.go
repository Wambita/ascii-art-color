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
		fmt.Println("Usage: go run . [OPTION] and/or [BANNER] [STRING] \n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
		return
	}
	var (
		inputString    string
		lettersToColor string
		color          string
		banner         string
		fileName       string
	)
	// case: no colors selected(input string only)
	if len(os.Args) == 2 {
		if !(strings.HasPrefix(os.Args[1], "--color=")) {
			inputString = os.Args[1]
			lettersToColor = inputString
			color = "resetCode"
			fileName = "standard.txt"
		} else {
			fmt.Println("Usage: go run . [OPTION] and/or [BANNER] [STRING] \n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
			return
		}
	}
	// case color selected but no specified letters to be colored (color+string)
	if len(os.Args) == 3 {
		if strings.HasPrefix(os.Args[1], "--color=") {
			colorFlag := os.Args[1]
			inputString = os.Args[2]
			fileName = "standard.txt"
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
	}
	// case color selected , letters to be colored specified and input string (color+letter+string)
	if len(os.Args) == 4 {
		colorFlag := os.Args[1]
		inputString = os.Args[3]
		lettersToColor = os.Args[2]
		fileName = "standard.txt"
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
	// case string+banner
	if len(os.Args) == 3 && ((os.Args[2] == "shadow") || (os.Args[2] == "thinkertoy") || (os.Args[2] == "standard")) {
		inputString = os.Args[1]
		lettersToColor = inputString
		color = "resetCode"

		banner = os.Args[2]

		switch banner {
		case "shadow":
			fileName = "shadow.txt"
		case "thinkertoy":
			fileName = "thinkertoy.txt"
		case "standard":
			fileName = "standard.txt"
		default:
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
			return
		}

	}
	// case color+string+banner
	if len(os.Args) == 4 && ((os.Args[3] == "shadow") || (os.Args[3] == "thinkertoy") || (os.Args[3] == "standard")) {

		colorFlag := os.Args[1]
		inputString = os.Args[2]
		lettersToColor = inputString
		banner = os.Args[3]

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

		switch banner {
		case "shadow":
			fileName = "shadow.txt"
		case "thinkertoy":
			fileName = "thinkertoy.txt"
		case "standard":
			fileName = "standard.txt"
		default:
			fmt.Println("Usage: go run . [OPTION] [STRING] \n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
			return
		}
	}

	// case  color+lettertocolor+string+banner
	if len(os.Args) == 5 && ((os.Args[4] == "shadow") || (os.Args[4] == "thinkertoy") || (os.Args[4] == "standard")) {

		inputString = os.Args[3]
		lettersToColor = os.Args[2]
		colorFlag := os.Args[1]
		banner = os.Args[4]

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

		switch banner {
		case "shadow":
			fileName = "shadow.txt"
		case "thinkertoy":
			fileName = "thinkertoy.txt"
		case "standard":
			fileName = "standard.txt"
		default:
			fmt.Println("Usage: go run . [OPTION] [STRING] \n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
			return
		}
	}

	// Extract the letters to be colored and the rest of the string
	// lettersToColor, inputString = extractLettersToColor(args)

	characterMap, err := asciiart.CreateMap(fileName)
	if err != nil {
		fmt.Println("Error reading map:", err)
		return
	}

	asciiart.DisplayAsciiArtWithPartialColor(characterMap, inputString, lettersToColor, color)
}
