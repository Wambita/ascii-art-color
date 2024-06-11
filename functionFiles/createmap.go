package asciiart

import (
	"bufio"
	"fmt"
	"hash/crc32"
	"os"
	"path/filepath"
)

func CreateMap(fileName string) (map[rune][]string, error) {
	// Open and read a file specified by the given file path(s), creating an ASCII art map.
	// Check if the file exists
	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist:", fileName)
		} else {
			fmt.Println("Error checking file status:", err)
		}
		return nil, err
	}
	defer file.Close()
	// Check for empty file
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	if stat.Size() == 0 {
		return nil, fmt.Errorf("error: Character map file '%s' is empty", fileName)
	}
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	var data []byte
	for _, value := range lines {
		for _, value2 := range value {
			data = append(data, byte(value2))
		}
	}
	crc32Table := crc32.MakeTable(crc32.IEEE)
	checksum := crc32.Checksum(data, crc32Table)
	if !(checksum == 0x9ffd59bc || checksum == 0x2f465361 || checksum == 0x6ee86a07) {
		return nil, fmt.Errorf("file modified, undo changes made to %q ", fileName)
	}
	err = scanner.Err()
	if err != nil {
		return nil, err
	}
	// Validate the file path and checks for empty files or non-text file extensions.
	if filepath.Ext(fileName) != ".txt" {
		fmt.Println("Wrong extension, use .txt")
		return nil, err
	}
	// Parse the file line by line, constructing the ASCII art map.
	// Uses a map where the key is the ASCII character (rune) and the value is a slice of strings representing the lines of the ASCII art.
	characterMap := make(map[rune][]string)
	var currentChar rune = ' '
	for _, line := range lines {
		if len(line) == 0 { // Skip empty lines
			continue
		}
		if len(characterMap[currentChar]) == 8 {
			currentChar++
			characterMap[currentChar] = []string{}
		}
		characterMap[currentChar] = append(characterMap[currentChar], line)
	}
	// Returns the constructed map or an error if any occurs during file operations or map creation.
	return characterMap, nil
}
