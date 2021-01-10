package utils

import (
	"fmt"
	"log"
	"net/http"
)

// CheckupResults ... Struct which holds a Checkup results
type CheckupResults struct {
	Endpoint string
	Code     int
	Result   int
	Pass     bool
}

// Checkup ... utility function to run one test
func Checkup(client *http.Client, code int, endpoint string) CheckupResults {
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
	results := CheckupResults{
		Endpoint: endpoint,
		Code:     code,
		Result:   resp.StatusCode,
		Pass:     resp.StatusCode == code,
	}
	return results
}
