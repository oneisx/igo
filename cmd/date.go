package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "igo/util"
    "strconv"
    "time"
)

var dateOriginUsageFunc string

var dateCmd = &cobra.Command{
    Use:   "date",
    Short: "transfer timestamp to date",
    Long:  `transfer timestamp to date`,
    Args: func(cmd *cobra.Command, args []string) error {
        return validTimestamp(args)
    },
    Run: func(cmd *cobra.Command, args []string) {
        doTransferTimestamp(args[0])
    },
}

func doTransferTimestamp(timestamp string) {
    i, err := strconv.ParseInt(timestamp, 10, 64)
    if err != nil {
        _ = fmt.Errorf("timestamp is invalid")
        return
    }
    tLen := len(timestamp)
    var tm time.Time
    if tLen == 10 {
        tm = time.Unix(i, 0)
    } else {
        tm = time.UnixMilli(i)
    }
    fmt.Println(tm)
    util.WriteClipboard(tm.String())
}

func validTimestamp(args []string) error {
    length := len(args)
    if length != 1 {
        return fmt.Errorf("accepts %d arg(s), received %d", 1, length)
    }

    tLen := len(args[0])
    if tLen != 10 && tLen != 13 {
        return fmt.Errorf("length of timestamp is invalid, only support second and millisecond, length is %d", tLen)
    }
    return nil
}

func init() {
    rootCmd.AddCommand(dateCmd)
    dateOriginUsageFunc = dateCmd.UsageString()
    dateCmd.SetUsageFunc(dateUsageFunc)
}

var dateUsageFunc = func(cmd *cobra.Command) error {
    fmt.Println(dateOriginUsageFunc)
    fmt.Println(`Examples:
  Non-interactive:
    igo date 1639238044
    igo date 1639188919040
  Interactive:
    date 1639238044
    date 1639188919040`)
    return nil
}
