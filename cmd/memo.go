package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var memoCmd = &cobra.Command{
    Use:   "memo",
    Short: "memo",
    Long:  `memo`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("memo called")
    },
}

func init() {
    rootCmd.AddCommand(memoCmd)
}
