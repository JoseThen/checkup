package cmd

import (
	"fmt"
	"os"

	utils "github.com/JoseThen/checkup/utils"
	tea "github.com/charmbracelet/bubbletea"
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
		checkupList := []utils.CheckupResults{}

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
				httpClient.CloseIdleConnections()
				checkupList = append(checkupList, checkup)
				if !checkup.Pass {
					exitCode = 3
				}
			}
		}

		m := utils.TeaTable(checkupList)
		program := tea.NewProgram(m)
		if _, err := program.Run(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
		if exitCode == 0 {
			defer os.Exit(exitCode)
		} else {
			defer utils.CustomErrorOut("**exam failed**", exitCode)
		}
	},
}

func init() {
	rootCmd.AddCommand(examCmd)
	examCmd.Flags().StringVarP(&File, "file", "f", "", "File to run exam against")
	examCmd.MarkFlagRequired("file")
}
