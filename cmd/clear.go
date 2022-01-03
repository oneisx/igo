package cmd

import (
    "igo/util"

    "github.com/spf13/cobra"
)

// since v1.0.0
var clearCmd = &cobra.Command{
    Use:   "clear",
    Short: "clear screen",
    Long:  `clear screen`,
    Run: func(cmd *cobra.Command, args []string) {
        util.Operator.ClearScreen()
    },
}

func init() {
    rootCmd.AddCommand(clearCmd)
}
