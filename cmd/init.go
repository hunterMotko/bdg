/*
Copyright © 2024 Hunter Motko <huntermotko.dev@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/hunterMotko/budgot/internal/database"
	"github.com/hunterMotko/budgot/internal/views"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initalize your expenses and income",
	Long: `Initalize the main information that you will need to base your budget

    Initalize your config file
    What your starting balance
    What your planned Expenses
    What your planned Income
  `,
	Run: func(cmd *cobra.Command, args []string) {
		CreateSqliteFile()
		service := database.New(dbPath)
		res := service.Init()
    message := res["message"]
		if  message != "success" {
      log.Fatalf("INIT DB ERROR: %s", message)
		}
		data, err := views.RunInit()
		if err != nil {
			log.Fatal("INIT FORM ERROR")
		}
		res = service.AddPlanned(data)
    message = res["message"]
		if message != "success" {
      log.Fatalf("INSERT ERROR: %s", message)
		}
		fmt.Println(
			lipgloss.NewStyle().
				Width(40).
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("63")).
				Padding(1, 2).
				Render("\n\tInitialize Successful\n"),
		)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func CreateSqliteFile() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("CONFIG DIR ERROR: %v", err)
	}
	dirPath := fmt.Sprintf("%s/%s", home, os.Getenv("DB_DIR"))
	err = os.Mkdir(dirPath, 744)
	if err != nil {
		log.Fatalf("DIR WRITE ERROR: %v", err)
	}
}
