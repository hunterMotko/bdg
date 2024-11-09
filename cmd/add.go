/* Copyright © 2024 hdm <huntermotko@gmail.com> */
package cmd

import (
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [command]",
	Short: "Add an Income or Expense Transaction",
	Long:  ` `,
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.AddCommand(addExpense)
	addCmd.AddCommand(addIncome)
}
