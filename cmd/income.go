/*
Copyright © 2024 Hunter Motko<huntermotko.dev@gmail.com>
*/
package cmd

import (
	"log"

	"github.com/hunterMotko/bdg/internal/database"
	"github.com/hunterMotko/bdg/internal/utils"
	"github.com/hunterMotko/bdg/internal/views"
	"github.com/spf13/cobra"
)

// incomeCmd represents the income command
var incomeCmd = &cobra.Command{
	Use:   "income",
	Short: "Show current planned, actual, and diff of Income for the current month",
	Long: `
  `,
	Run: func(cmd *cobra.Command, args []string) {
		if !utils.CheckDBFileExists(conf) {
			log.Fatal("DB File does not exist")
		}
    service := database.New(conf.String())
    pl, err := service.GetPlannedIncome()
    if err != nil {
		  log.Fatalf("GET PLANNED ERROR: %v", err)
    }
    mp, err := service.GetIncomeRecords()
    if err != nil {
		  log.Fatalf("GET INCOME RECORDS ERROR: %v", err)
    }
    views.GetTable(pl, mp)
	},
}

func init() {
	rootCmd.AddCommand(incomeCmd)
}
