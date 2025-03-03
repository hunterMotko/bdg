/*
Copyright © 2025 Hunter Motko <huntermotko.dev@gmail.com>
*/
package cmd

import (
	"log"

	"github.com/hunterMotko/budgot/internal/database"
	"github.com/hunterMotko/budgot/internal/utils"
	"github.com/hunterMotko/budgot/internal/views"
	"github.com/spf13/cobra"
)

// sumCmd represents the sum command
var sumCmd = &cobra.Command{
	Use:   "sum",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !utils.CheckDBFileExists(conf) {
			log.Fatal("DB File does not exist")
		}
		service := database.New(conf.String())
		service.Health()
		sums, err := service.GetSums()
		if err != nil {
			log.Fatal(err)
		}
		views.RunSum(sums)
	},
}

func init() {
	rootCmd.AddCommand(sumCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sumCmd.PersistentFlags().String("foo", "", "A help for foo")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sumCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

