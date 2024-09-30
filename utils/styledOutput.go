package utils

import (
	"os"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

const (
	purple    = lipgloss.Color("99")
	gray      = lipgloss.Color("245")
	lightGray = lipgloss.Color("241")
)

func Show(checkups []CheckupResults) *table.Table {
	re := lipgloss.NewRenderer(os.Stdout)

	var (
		// HeaderStyle is the lipgloss style used for the table headers.
		HeaderStyle = re.NewStyle().Foreground(purple).Bold(true).Align(lipgloss.Center)
		// CellStyle is the base lipgloss style used for the table rows.
		CellStyle = re.NewStyle().Padding(0, 1)
		// OddRowStyle is the lipgloss style used for odd-numbered table rows.
		OddRowStyle = CellStyle.Foreground(gray)
		// EvenRowStyle is the lipgloss style used for even-numbered table rows.
		EvenRowStyle = CellStyle.Foreground(lightGray)
		// BorderStyle is the lipgloss style used for the table border.
		BorderStyle = lipgloss.NewStyle().Foreground(purple)
	)
	t := table.New()
	t.Headers("Endpoint", "Code", "Result", "Latency", "Pass")

	// Table Styling happening here
	t.Border(lipgloss.NormalBorder())
	t.BorderStyle(BorderStyle)
	t.StyleFunc(
		func(row, col int) lipgloss.Style {
			switch {
			case row == 0:
				return HeaderStyle
			case row%2 == 0:
				return EvenRowStyle
			default:
				return OddRowStyle
			}
		})

	for count := range checkups {
		t.Row(
			checkups[count].Endpoint,
			strconv.Itoa(checkups[count].Code),
			strconv.Itoa(checkups[count].Result),
			strconv.FormatInt(checkups[count].Lantency.Milliseconds(), 10),
			strconv.FormatBool(checkups[count].Pass),
		)
	}

	return t
}
