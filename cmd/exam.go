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

// File ... Variable for File flag
var File string

// examCmd represents the exam command
var examCmd = &cobra.Command{
	Use:   "exam",
	Short: "Run a checkup against a file.",
	Long:  `Runs checkup against a file`,
	Run: func(cmd *cobra.Command, args []string) {
		// Flag Parsing
		file, _ := cmd.Flags().GetString("file")
		auth, _ := cmd.Flags().GetBool("auth")
		exam := utils.ReadExam(file)
		exitCode := 0
		// Setup http Client
		var httpClient = &http.Client{
			Timeout: time.Second * 10,
		}
		w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
		fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t", "Endpoint", "Code", "Result", "Pass")
		fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t", "--------", "----", "------", "----")
		for _, test := range exam.Tests {
			for _, path := range test.Paths {
				// Setup check Request with above variables
				checkRequest := &utils.CheckupRequest{
					Client:   httpClient,
					Code:     test.Code,
					Endpoint: exam.Endpoint + path,
					Auth:     auth,
				}
				checkup := utils.Checkup(*checkRequest)
				fmt.Fprintf(w, "\n %s\t%d\t%d\t%v\t", checkup.Endpoint, checkup.Code, checkup.Result, checkup.Pass)
				if checkup.Pass == false {
					exitCode = 1
				}
			}
		}
		w.Flush()
		fmt.Println()
		defer os.Exit(exitCode)
	},
}

func init() {
	rootCmd.AddCommand(examCmd)
	examCmd.Flags().StringVarP(&File, "file", "f", "", "File to run exam against")
	examCmd.MarkFlagRequired("file")
}
