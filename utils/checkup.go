package utils

import (
	"fmt"
	"log"
	"net/http"
)

// Checkup ... utility function to run one test
func Checkup(client *http.Client, code int, endpoint string) bool {
	request, err := http.NewRequest("GET", endpoint, nil)
	request.Header.Set("User-Agent", "CheckupCli/1.0")
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Printf("error in checkup function: %v", err)
	}
	defer resp.Body.Close()
	return resp.StatusCode == code
}
