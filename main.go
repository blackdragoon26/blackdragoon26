package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	ReadmeFile  = "README.md"
	StartMarker = ""
	EndMarker   = ""
	ApiURL      = "https://icanhazdadjoke.com/"
)

type JokeResponse struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func main() {
	// 1. Fetch the Joke
	client := &http.Client{}
	req, err := http.NewRequest("GET", ApiURL, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "GitHub Readme Bot (https://github.com/blackdragoon26/blackdragoon26)")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	var data JokeResponse
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		panic(err)
	}

	fmt.Println("Fetched joke:", data.Joke)

	// 2. Read README
	contentBytes, err := os.ReadFile(ReadmeFile)
	if err != nil {
		panic(err)
	}
	content := string(contentBytes)

	// 3. Find Markers (Robust Way)
	startIndex := strings.Index(content, StartMarker)
	if startIndex == -1 {
		fmt.Println("Start marker not found in README.md")
		os.Exit(1)
	}

	// Look for the End marker ONLY after the Start marker
	restOfContent := content[startIndex:]
	endIndexOffset := strings.Index(restOfContent, EndMarker)
	if endIndexOffset == -1 {
		fmt.Println("End marker not found in README.md")
		os.Exit(1)
	}
	endIndex := startIndex + endIndexOffset

	// 4. Construct New Content
	newSection := fmt.Sprintf("%s\n### Worst Dad Joke of the day ￣\\_(ツ)_/￣ \n> %s\n%s", StartMarker, data.Joke, EndMarker)

	// Replace the old part with the new part
	newContent := content[:startIndex] + newSection + content[endIndex+len(EndMarker):]

	// 5. Write Back
	err = os.WriteFile(ReadmeFile, []byte(newContent), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("README updated successfully.")
}
