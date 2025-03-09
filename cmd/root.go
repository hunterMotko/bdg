/*
Copyright © 2024 hdm <huntermotko.dev@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/hunterMotko/bdg/internal/config"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/cobra"
)

var (
	dbPath string
  conf *config.Config
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "budgot",
	Short: "My Personal budgeting tool",
	Long: `
  How to use this personal budgeting tool: 
    Init: Initialize you starting balance and what your planned income and expenses should be
    Add: Add income/expense records as you recieve them to view how your budgeting is going
    Income: View your Income planned - actual - diff 
    Expenses: View your Expenses planned - actual - diff 
  `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(os.Stdout, "%s", cmd.UsageString())
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
  rootCmd.PersistentFlags().StringVarP(&dbPath, "config-path", "c", ".config/bdg", "Config Path string")
  home, err := os.UserHomeDir()
  if err != nil {
    fmt.Fprintf(os.Stderr, "ERROR finding home dir: %v\n", err)
  }
  conf = config.NewConfig(home, dbPath, "bdg.sqlite")
}
