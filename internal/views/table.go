package views

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/hunterMotko/bdg/internal/data"
)

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
      // TODO: Do I want them to be able to update a record here?
      // if so do I allow this to update plannned? Or should planned be update specificly and every month? 
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.table.SelectedRow()[1]),
			)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func setTable(rows []table.Row) model {
	columns := []table.Column{
		{Title: "Category", Width: 20},
		{Title: "Planned", Width: 20},
		{Title: "Actual", Width: 20},
		{Title: "Diff", Width: 20},
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

	return model{t}
}

func GetTable(planned []data.Category, recs map[int]float64) {
	rows := mergeData(planned, recs)
	m := setTable(rows)
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("ERROR RUNNING PROGRAM: ", err)
		os.Exit(1)
	}
}

func mergeData(planned []data.Category, recs map[int]float64) []table.Row {
	rows := []table.Row{}
	for _, p := range planned {
		sub := p.Planned - recs[p.Id]
		temp := []string{
      p.Name, 
      fmt.Sprintf("%.2f", p.Planned), 
      fmt.Sprintf("%.2f", recs[p.Id]), 
      fmt.Sprintf("%.2f", sub),
    }
		rows = append(rows, temp)
	}
	return rows
}
