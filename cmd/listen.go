package cmd

import (
	"fmt"
	"os"
	"strconv"

	utils "github.com/JoseThen/checkup/utils"
	"github.com/spf13/cobra"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	// BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, tea.Quit
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

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

		columns := []table.Column{
			{Title: "Endpoint", Width: 50},
			{Title: "Code", Width: 5},
			{Title: "Result", Width: 10},
			{Title: "Latency", Width: 10},
			{Title: "Pass", Width: 5},
		}

		rows := []table.Row{
			{
				checkup.Endpoint,
				strconv.Itoa(checkup.Code),
				strconv.Itoa(checkup.Result),
				strconv.FormatInt(checkup.Lantency.Milliseconds(), 10),
				strconv.FormatBool(checkup.Pass)},
		}

		t := table.New(
			table.WithColumns(columns),
			table.WithRows(rows),
			table.WithFocused(true),
			table.WithHeight(7),
		)

		s := table.DefaultStyles()
		s.Header = s.Header.
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240")).
			BorderBottom(true).
			Bold(false)
		s.Selected = s.Selected.
			Foreground(lipgloss.Color("229")).
			Background(lipgloss.Color("57")).
			Bold(false)
		t.SetStyles(s)

		m := model{t}
		if _, err := tea.NewProgram(m).Run(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}

		// w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
		// fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t", "Endpoint", "Code", "Result", "Latency", "Pass")
		// fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t", "--------", "----", "------", "-------", "----")
		// fmt.Fprintf(w, "\n %s\t%d\t%d\t%dms\t%v\t", checkup.Endpoint, checkup.Code, checkup.Result, checkup.Lantency.Milliseconds(), checkup.Pass)
		// w.Flush()
		// fmt.Println()
		if checkup.Pass {
			defer os.Exit(0)
		} else {
			defer utils.CustomErrorOut("\n**checkup failed**", 2)
		}

	},
}

func init() {
	rootCmd.AddCommand(listenCmd)
	listenCmd.Flags().IntVarP(&Code, "code", "c", 200, "Expected status code")
	listenCmd.Flags().StringVarP(&Endpoint, "endpoint", "e", "", "Endpoint for checkup")
	listenCmd.MarkFlagRequired("endpoint")
}
