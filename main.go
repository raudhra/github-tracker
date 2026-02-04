package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type track struct {
	Type string `json:"type"`

	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`

	Payload struct {
		Commits []struct{} `json:"commits"`
	} `json:"payload"`
}

func main() {
	var username string
	token := os.Getenv("GITHUB_TOKEN")
	fmt.Print("Enter The Github Username: ")
	fmt.Scan(&username)
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error Sending Request Please Check Username")
	}
	req.Header.Set("User-Agent", "raudhra")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-Github-Api-Version", "2022-11-28")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error Sending Request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response Body")
	}
	var Track []track
	if err := json.Unmarshal(respBody, &Track); err != nil {
		fmt.Printf("Error")
	}
	//fmt.Printf("Output:\n %+v\n", Track)

	for _, e := range Track {
		switch e.Type {
		case "PushEvent":
			fmt.Printf("- Pushed %d commits to %s\n", len(e.Payload.Commits), e.Repo.Name)
		case "WatchEvent":
			fmt.Printf("- Starred %s\n", e.Repo.Name)
		case "IssuesEvent":
			fmt.Printf("- Opened a new issue in %s\n", e.Repo.Name)
		case "ForkEvent":
			fmt.Printf("- Forked %s\n", e.Repo.Name)
		}

	}
}
