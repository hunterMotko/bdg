package cmd

import (
	"fmt"
	"log"

	"github.com/charmbracelet/lipgloss"
	"github.com/hunterMotko/bdg/internal/database"
	"github.com/hunterMotko/bdg/internal/utils"
	"github.com/hunterMotko/bdg/internal/views"
	"github.com/spf13/cobra"
)

var addExpense = &cobra.Command{
	Use:   "expense",
	Short: "Add a Expense Record",
	Long:  ``,
	Run:   runExpense,
}

func runExpense(cmd *cobra.Command, args []string) {
  if !utils.CheckDBFileExists(conf) {
    log.Fatalf("Check file path or use INIT command: %s", conf.String())
  }
	db := database.New(conf.String())
	opts := []string{
		"food", "gifts", "medical", "home", "transportation", "personal", "pets", "utilities", "travel", "debt", "other",
	}
	rec, err := views.RunAdd("Expense", opts)
	if err != nil {
		log.Fatalf("ADD ERROR: %v", err)
	}
	res := db.InsertExpense(rec)
	if res["message"] != "success" {
		log.Fatalf("INSERT ERROR: %v", err)
	}
	fmt.Println(
		lipgloss.NewStyle().
			Width(40).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")).
			Padding(1, 2).
			Render("\n\tExpense Record Added\n"),
	)
}
