package cmd

import (
	"fmt"
	"net/http"
	"os"
	"text/tabwriter"
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
		checkup := utils.Checkup(httpClient, code, endpoint)
		w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
		fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t", "Endpoint", "Code", "Result", "Pass")
		fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t", "--------", "----", "------", "----")
		fmt.Fprintf(w, "\n %s\t%d\t%d\t%v\t", checkup.Endpoint, checkup.Code, checkup.Result, checkup.Pass)
		w.Flush()
		fmt.Println()
		if checkup.Pass {
			defer os.Exit(0)
		} else {
			defer os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(listenCmd)
	listenCmd.Flags().IntVarP(&Code, "code", "c", 200, "Expected status code")
	listenCmd.Flags().StringVarP(&Endpoint, "endpoint", "e", "", "Endpoint for checkup")
	listenCmd.MarkFlagRequired("endpoint")
}
