package main

import (
	"encoding/json"
	"io"
	"jira-get-tickets/excel"
	"jira-get-tickets/structs"
	"log"
	"net/http"
)

func main() {

	// Example public URI
	response, err := http.Get("https://jira.atlassian.com/rest/api/latest/search?jql=project%20=%20JRASERVER%20AND%20status%20=%20%22In%20Progress%22%20AND%20component%20=%20Accessibility")

	if err != nil {
		log.Fatal(err)
	}

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var tickets structs.TicketsData

	err = json.Unmarshal(responseData, &tickets)

	if err != nil {
		log.Fatal(err)
	}

	if err = excel.ConvertToExcel(&tickets); err != nil {
		log.Fatal(err)
	}
}
