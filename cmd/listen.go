package cmd

import (
	"fmt"
	"os"

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
		// Flag parsing
		code, _ := cmd.Flags().GetInt("code")
		auth, _ := cmd.Flags().GetBool("auth")
		endpoint, _ := cmd.Flags().GetString("endpoint")

		// Setup check Request with above variables
		checkRequest := &utils.CheckupRequest{
			Client:   httpClient,
			Code:     code,
			Endpoint: endpoint,
			Auth:     auth,
		}
		checkup := utils.Checkup(*checkRequest)
		httpClient.CloseIdleConnections()

		checkupList := []utils.CheckupResults{
			checkup,
		}

		table := utils.Show(checkupList)
		fmt.Println(table)

		if checkup.Pass {
			defer os.Exit(0)
		} else {
			defer utils.CustomErrorOut("**checkup failed**", 2)
		}

	},
}

func init() {
	rootCmd.AddCommand(listenCmd)
	listenCmd.Flags().IntVarP(&Code, "code", "c", 200, "Expected status code")
	listenCmd.Flags().StringVarP(&Endpoint, "endpoint", "e", "", "Endpoint for checkup")
	listenCmd.MarkFlagRequired("endpoint")
}
