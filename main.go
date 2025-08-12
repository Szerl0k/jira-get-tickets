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

	uri := "https://jira.atlassian.com/rest/api/latest/search?jql=project%20=%20JRASERVER%20AND%20status%20=%20%22In%20Progress%22%20AND%20component%20in%20(Accessibility,%20%22Administration%20-%20S3%20Object%20Storage%22,%20%22Environment%20-%20Database%22)"

	responseData, err := sendHttpRequest(uri)

	var tickets structs.TicketsData

	err = json.Unmarshal(responseData, &tickets)

	if err != nil {
		log.Fatal(err)
	}

	if err = excel.ConvertToExcel(&tickets); err != nil {
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
