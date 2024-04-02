package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

func ListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all commands that are available",
		Run: func(_ *cobra.Command, _ []string) {
			listCommand()
		},
		Aliases: []string{"l", "lis", "ls"},
	}

	return cmd
}

//func listCommand() {
//	table := tablewriter.NewWriter(os.Stdout)
//	table.SetHeader([]string{"command", "description"})
//	table.SetRowLine(true)
//
//	for k, v := range unmarshalIndex() {
//		table.Append([]string{k, v.Description})
//	}
//
//	table.Render()
//}

// TODO: split the window vertically and output the content

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				tea.Println(m.table.SelectedRow()[0]))
		}
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func listCommand() {
	columns := []table.Column{
		{Title: "command", Width: 30},
		{Title: "description", Width: 80},
	}

	// load data
	data := unmarshalIndex()
	rows := make([]table.Row, 0, len(data))
	for k, v := range data {
		rows = append(rows, table.Row{k, v.Description})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(14), // Set how many items be displayed at the most in table.
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
}
