package cmd

import (
	"fmt"
	"net/http"
	"time"

	utils "github.com/JoseThen/checkup/utils"
	"github.com/spf13/cobra"
)

// Code ...Variable for code flag
var Code int

// Endpoint ... Variable for endpoint flag
var Endpoint string

// listenCmd represents the listen command
var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "Ensure endpoint resolves with correct status code",
	Long:  `Send a request to a given endpoint and assert that the status code matches your expected status code.`,
	Run: func(cmd *cobra.Command, args []string) {
		code, _ := cmd.Flags().GetInt("code")
		endpoint, _ := cmd.Flags().GetString("endpoint")
		var httpClient = &http.Client{
			Timeout: time.Second * 10,
		}
		fmt.Printf("listen called with Code: %v and Endpoint: %v\n", code, endpoint)
		valid := utils.Checkup(httpClient, code, endpoint)
		if valid {
			fmt.Println("Good Test")
		} else {
			fmt.Println("Bad Test")
		}

	},
}

func init() {
	rootCmd.AddCommand(listenCmd)
	listenCmd.Flags().IntVarP(&Code, "code", "c", 200, "Expected status code")
	listenCmd.Flags().StringVarP(&Endpoint, "endpoint", "e", "", "Endpoint for checkup")
	listenCmd.MarkFlagRequired("endpoint")
}
