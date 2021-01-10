package cmd

import (
	"fmt"
	"net/http"
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
		file, _ := cmd.Flags().GetString("file")
		pass := true
		exam := utils.ReadYaml(file)
		var httpClient = &http.Client{
			Timeout: time.Second * 10,
		}
		for _, test := range exam.Tests {
			for _, path := range test.Paths {
				checkup := utils.Checkup(httpClient, test.Code, exam.Endpoint+path)
				if checkup == false {
					pass = false
				}
			}
		}
		if pass {
			fmt.Println("Good Test")
		} else {
			fmt.Println("Bad Test")
		}
	},
}

func init() {
	rootCmd.AddCommand(examCmd)
	examCmd.Flags().StringVarP(&File, "file", "f", "", "File to run exam against")
	examCmd.MarkFlagRequired("file")
}
