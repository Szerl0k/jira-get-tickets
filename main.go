package main

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"io"
	"jira-get-tickets/excel"
	"jira-get-tickets/structs"
	"log"
	"net/http"
	"os"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	uri := os.Getenv("JIRA_URI")

	responseData, err := sendHttpRequest(uri)

	if err != nil {
		log.Fatal(err)
	}

	var tickets structs.TicketsData

	err = json.Unmarshal(responseData, &tickets)

	if err != nil {
		log.Fatal(err)
	}

	if err = excel.ExportTicketsToExcel(&tickets); err != nil {
		log.Fatal(err)
	}
}

func sendHttpRequest(uri string) ([]byte, error) {
	req, err := http.NewRequest("GET", uri, nil)

	if err != nil {
		return nil, err
	}

	err = godotenv.Load(".env")

	if err != nil {
		return nil, err
	}

	key := os.Getenv("JIRA_PAC")

	req.Header.Set("Authorization", "Bearer "+key)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return responseData, nil
}
