package cmd

import (
	"fmt"
	"github.com/spf13/cast"

	"github.com/spf13/cobra"
)

// dateCmd represents the date command
var dateCmd = &cobra.Command{
	Use:   "date",
	Short: "transfer timestamp to date string",
	Long: `transfer timestamp to date string`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("date called" + cast.ToString(args))
	},
}

func init() {
	rootCmd.AddCommand(dateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
