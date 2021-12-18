package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var memoCmd = &cobra.Command{
    Use:   "memo [WIP]",
    Short: "memo [WIP]",
    Long:  `memo [WIP]`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("memo called")
    },
}

func init() {
    rootCmd.AddCommand(memoCmd)
}
