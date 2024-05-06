package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize database",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := createDataDir()
		if err != nil {
			fmt.Println("Error initializing sheet", err)
		}
		fmt.Println("Initialized sheet")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
