package utils

import (
	"encoding/base64"
	"net/http"
	"os"
)

// CheckupRequest ... Struct holding all data needed for a checkup
type CheckupRequest struct {
	Client   *http.Client
	Code     int
	Endpoint string
	Auth     bool
}

// CheckupResults ... Struct which holds a Checkup results
type CheckupResults struct {
	Endpoint string
	Code     int
	Result   int
	Pass     bool
}

// Checkup ... utility function to run one test
func Checkup(healthForm CheckupRequest) CheckupResults {
	request, err := http.NewRequest("GET", healthForm.Endpoint, nil)
	ErrorCheck(err)

	if request.URL.Scheme == "" {
		ErrorOut("You must use a supported protocol scheme like http or https")
	}

	request.Header.Set("User-Agent", "CheckupCli/1.0")
	if healthForm.Auth {
		request.Header.Add("Authorization", "Basic "+addBasicAuth(os.Getenv("CU_USER"), os.Getenv("CU_PASS")))
		healthForm.Client.CheckRedirect = addAuthOnRedirect
	}

	resp, err := healthForm.Client.Do(request)
	if err != nil {
		ErrorCheck(err)
	}
	defer resp.Body.Close()

	results := CheckupResults{
		Endpoint: healthForm.Endpoint,
		Code:     healthForm.Code,
		Result:   resp.StatusCode,
		Pass:     resp.StatusCode == healthForm.Code,
	}

	return results
}

func addBasicAuth(user, pass string) string {
	auth := user + ":" + pass
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// We need this function because on a redirect with net/http you lose your items in your header
func addAuthOnRedirect(req *http.Request, via []*http.Request) error {
	req.Header.Add("Authorization", "Basic "+addBasicAuth(os.Getenv("CU_USER"), os.Getenv("CU_PASS")))
	return nil
}
