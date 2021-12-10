package cmd

import (
	"igo/util"

	"github.com/spf13/cobra"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "clear screen",
	Long: `clear screen`,
	Run: func(cmd *cobra.Command, args []string) {
		util.ClearScreen()
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
