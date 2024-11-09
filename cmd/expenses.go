/*
Copyright © 2024 Hunter Motko<huntermotko.dev@gmail.com>
*/
package cmd

import (
	"log"

	"github.com/hunterMotko/budgot/internal/database"
	"github.com/hunterMotko/budgot/internal/views"
	"github.com/spf13/cobra"
)

// expensesCmd represents the expenses command
var expensesCmd = &cobra.Command{
	Use:   "expenses",
	Short: "Show current planned, actual, and diff of Expenses",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		service := database.New(dbPath)
		pe, err := service.GetPlannedExpense()
		if err != nil {
			log.Fatalf("GET PLANNED ERROR: %v", err)
		}
		mp, err := service.GetExpenseRecords()
		if err != nil {
			log.Fatalf("GET INCOME RECORDS ERROR: %v", err)
		}
    views.GetTable(pe, mp)
	},
}

func init() {
	rootCmd.AddCommand(expensesCmd)
}
