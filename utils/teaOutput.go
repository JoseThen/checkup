package utils

import (
	"strconv"

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

func TeaTable(checkups []CheckupResults) model {
	columns := []table.Column{
		{Title: "Endpoint", Width: 50},
		{Title: "Code", Width: 5},
		{Title: "Result", Width: 10},
		{Title: "Latency", Width: 10},
		{Title: "Pass", Width: 5},
	}

	rows := []table.Row{}

	for count := range checkups {
		checkupRow := table.Row{
			checkups[count].Endpoint,
			strconv.Itoa(checkups[count].Code),
			strconv.Itoa(checkups[count].Result),
			strconv.FormatInt(checkups[count].Lantency.Milliseconds(), 10),
			strconv.FormatBool(checkups[count].Pass),
		}
		rows = append(rows, checkupRow)
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
		Bold(true)
	// We setup selected styles because it's used by default
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("fff")).
		// Background(lipgloss.Color("0")).
		Bold(false)
	s.Cell = s.Cell.
		Foreground(lipgloss.Color("fff")).
		// Background(lipgloss.Color("0")).
		Bold(false)
	t.SetStyles(s)

	m := model{t}
	return m
}
