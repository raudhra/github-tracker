package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type track struct {
	Event string `json:"type"`
}

func main() {
	var username string
	//token := ""
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
		fmt.Errorf("Error reading response Body")
	}
	var Track []track
	if err := json.Unmarshal(respBody, &Track); err != nil {
		fmt.Printf("Error")
	}
	fmt.Printf("Output:\n %+v\n", Track)

}
