package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"jira-get-tickets/excel"
	"jira-get-tickets/structs"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	uri := os.Getenv("JIRA_URI")

	responseData, err := sendHttpRequest(uri)

	if err != nil {
		logToFile(err.Error())
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

		if err == io.EOF || strings.Contains(err.Error(), "EOF") {
			return nil, fmt.Errorf("EOF error: server closed connection unexpectedly")
		}

		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusRequestTimeout || resp.StatusCode == http.StatusGatewayTimeout {
		return nil, fmt.Errorf("server responded with timeout: %d, %s", resp.StatusCode, resp.Status)
	}

	if resp.StatusCode >= 500 {
		return nil, fmt.Errorf("server error: %d, %s", resp.StatusCode, resp.Status)
	}

	responseData, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return responseData, nil
}

func logToFile(message string) {
	f, err := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)
	logger.Println(message)
}
