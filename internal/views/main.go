package views

import (
	"fmt"
	"math"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/hunterMotko/budgot/internal/data"
	"github.com/hunterMotko/budgot/internal/utils"
	"golang.org/x/term"
)

const (
	progressBarWidth  = 71
	progressFullChar  = "█"
	progressEmptyChar = "░"
	dotChar           = " • "
)

var (
	keywordStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	subtleStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	red          = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
	lightBlue    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12"))
	green        = lipgloss.NewStyle().Foreground(lipgloss.Color("35"))
	// ticksStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("79"))
	intStyle           = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	progressEmpty      = subtleStyle.Render(progressEmptyChar)
	dotStyle           = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(dotChar)
	mainStyle          = lipgloss.NewStyle().MarginLeft(2)
	roundedBorderGroup = lipgloss.NewStyle().Width(40).BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("63")).Padding(1, 2)
	// Gradient colors we'll use for the progress bar
	ramp = utils.MakeRampStyles("#B14FFF", "#00FFA3", progressBarWidth)
  docStyle = lipgloss.NewStyle().Margin(1, 2, 1, 2)
)

func RunMain(sums *data.Sums) {
	endBalance := sums.CalcEndBal()
	initialModel := MainModel{
		sums.Start,
		endBalance,
		sums.CalcPercChange(endBalance),
		sums.Saved(endBalance),
		sums.PlannedExpense,
		sums.TotalExpense,
		sums.ExpensePerc(),
		sums.PlannedIncome,
		sums.TotalIncome,
		sums.IncomePerc(),
		false,
	}

	p := tea.NewProgram(initialModel)
	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
	}
}

type MainModel struct {
	Start          float64
	End            float64
	Perc           float64
	Saved          float64
	PlannedExpense float64
	ActualExpense  float64
	ExPerc         float64
	PlannedIncome  float64
	ActualIncome   float64
	InPerc         float64
	Quitting       bool
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			m.Quitting = true
			return m, tea.Quit
		}
	}
	return m, nil
}
func (m MainModel) View() string {
	if m.Quitting {
		return "\n  KEEP BUDGETING!\n\n"
	}
	s := mainView(m)
	return mainStyle.Render("\n" + s + "\n\n")
}

func mainView(m MainModel) string {
  var doc strings.Builder
  width, height, _ := term.GetSize(int(os.Stdout.Fd()))
  fmt.Println(width, height)
  doc.WriteString(strings.Repeat(" ", width))

	var templ strings.Builder
	var group string
	if m.Perc < 0 {
		group = fmt.Sprintf(
			"\t-%f%%\n\tDescrease in total savings\n\t%s\n\t-$%2.f\n\tSaved this month\n",
			math.Abs(m.Perc)*100,
			strings.Repeat(dotStyle, 6),
			m.Saved,
		)
	} else {
		group = fmt.Sprintf(
			"\t+%f%%\n\tIncrease in total savings\n\t%s\n\t$%2.f\n\tSaved this month\n",
			math.Abs(m.Perc)*100,
			strings.Repeat(dotStyle, 6),
			m.Saved,
		)
	}
	templ.WriteString(lightBlue.Render("Monthly Budogt") + "\n\n")
	templ.WriteString(roundedBorderGroup.Render(group) + "\n\n")

	low, high := utils.BarPercentages(m.Perc)
	lowEx, highEx := utils.BarPercentages(m.ExPerc)
	lowIn, highIn := utils.BarPercentages(m.InPerc)

  // Main bar graph
	templ.WriteString(fmt.Sprintf("Start Balance: %s\n", intStyle.Render(fmt.Sprintf("%2.f", m.Start))))
	templ.WriteString(progressbar(low) + "\n")
	templ.WriteString(fmt.Sprintf("End Balance: %s\n", intStyle.Render(fmt.Sprintf("%2.f", m.End))))
	templ.WriteString(progressbar(high) + "\n\n")
  // Expenses bars
	templ.WriteString(lightBlue.Render("Expenses") + "\n")
	templ.WriteString("Planned: " + intStyle.Render(fmt.Sprintf("%2.f", m.PlannedExpense)) + "\n")
	templ.WriteString(progressbar(lowEx) + "\n")
	templ.WriteString("Actual: " + intStyle.Render(fmt.Sprintf("%2.f", m.ActualExpense)) + "\n")
	templ.WriteString(progressbar(highEx) + "\n\n")
  // Income bars
	templ.WriteString(lightBlue.Render("Income") + "\n")
	templ.WriteString("Planned: " + intStyle.Render(fmt.Sprintf("%2.f", m.PlannedIncome)) + "\n")
	templ.WriteString(progressbar(lowIn) + "\n")
	templ.WriteString("Actual: " + intStyle.Render(fmt.Sprintf("%2.f", m.ActualIncome)) + "\n")
	templ.WriteString(progressbar(highIn) + "\n\n")
	templ.WriteString(dotStyle + subtleStyle.Render("q, esc: quit") + dotStyle)

	return templ.String()
}

func progressbar(percent float64) string {
	w := float64(progressBarWidth)
	fullSize := int(math.Round(w * percent))
	var fullCells string
	for i := 0; i < fullSize; i++ {
		fullCells += ramp[i].Render(progressFullChar)
	}
	emptySize := int(w) - fullSize
	emptyCells := strings.Repeat(progressEmpty, emptySize)
	return fmt.Sprintf("%s%s", fullCells, emptyCells)
}
