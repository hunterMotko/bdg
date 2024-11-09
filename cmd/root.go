/*
Copyright © 2024 hdm <huntermotko.dev@gmail.com>
*/
package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/hunterMotko/budgot/internal/database"
	"github.com/hunterMotko/budgot/internal/views"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/cobra"
)

var (
	dbPath string
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
	Run: run,
}

func run(cmd *cobra.Command, args []string) {
	err := checkDbExists()
	if err != nil {
		log.Fatal(err)
	}
	service := database.New(dbPath)
	service.Health()
	sums, err := service.GetSums()
	if err != nil {
		log.Fatal(err)
	}
  views.RunMain(sums)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	dbPathString()
}

func dbPathString() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nHOME ENV ERROR: %v\n", err)
	}
	dbPath = fmt.Sprintf("%s/%s/%s", home, os.Getenv("DB_DIR"), os.Getenv("DB_FILE"))
}

func checkDbExists() error {
	if _, err := os.Stat(dbPath); errors.Is(err, os.ErrNotExist) {
		return errors.New("Please Use Init Command before using")
	}
	return nil
}
