package cmd

import (
	"fmt"

	"github.com/karstenpedersen/sheet/db"
	"github.com/karstenpedersen/sheet/models"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add entry",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		title, err := cmd.Flags().GetString("title")
		if err != nil {
			fmt.Println("Provide a title")
			return
		}
		shortcut, err := cmd.Flags().GetString("shortcut")
		if err != nil {
			fmt.Println("Provide a shortcut")
			return
		}
		description, _ := cmd.Flags().GetString("description")
		scope, _ := cmd.Flags().GetString("scope")

		sc := models.NewShortcut(title, description, scope, shortcut)
		err = db.InsertShortcut(sc)
		if err != nil {
			fmt.Println("Error adding shortcut:", err)
		}
	},
}

func init() {
	addCmd.Flags().StringP("title", "t", "", "")
	addCmd.Flags().StringP("shortcut", "s", "", "")
	addCmd.Flags().StringP("desc", "d", "", "")
	addCmd.Flags().StringP("scope", "c", "", "")
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
