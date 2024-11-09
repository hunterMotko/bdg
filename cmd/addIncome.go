package cmd

import (
	"fmt"
	"log"

	"github.com/charmbracelet/lipgloss"
	"github.com/hunterMotko/budgot/internal/database"
	"github.com/hunterMotko/budgot/internal/views"
	"github.com/spf13/cobra"
)

var addIncome = &cobra.Command{
	Use:   "income",
	Short: "Add a Income Record",
	Long:  ``,
	Run:   runIncome,
}

func runIncome(cmd *cobra.Command, args []string) {
	db := database.New(dbPath)
	opts := []string{"savings", "paycheck", "bonus", "interest", "other"}
	rec, err := views.RunAdd("Income", opts)
	if err != nil {
		log.Fatalf("ADD ERROR: %v", err)
	}
	res := db.InsertIncome(rec)
	if res["message"] != "success" {
		log.Fatalf("INSERT ERROR: %v", err)
	}
	fmt.Println(
		lipgloss.NewStyle().
			Width(40).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")).
			Padding(1, 2).
			Render("\n\tIncome Record Added\n"),
	)
}
