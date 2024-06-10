package asciiart

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetFile(fileName string) {
	if !(fileName == "thinkertoy.txt" || fileName == "shadow.txt" || fileName == "standard.txt") {
		fmt.Println("Provide the correct banner file name <standard> or <shadow> or <thinkertoy>")
		return
	}
	fullURL := "https://learn.zone01kisumu.ke/git/root/public/raw/branch/master/subjects/ascii-art/" + fileName

	// create a blank file:
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error:", err)
	}

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			req.URL.Opaque = req.URL.Path
			return nil
		},
	}

	// put contents on file
	response, err := client.Get(fullURL)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer response.Body.Close()
	size, err := io.Copy(file, response.Body)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer file.Close()
	fmt.Printf("Downloaded a file %s with size %d\n Run the program again\n", fileName, size)
}
