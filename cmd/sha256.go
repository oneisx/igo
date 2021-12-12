package cmd

import (
    "fmt"
    "igo/util"

    "github.com/spf13/cobra"
)

var sha256US string
var sha256Key string

var sha256Cmd = &cobra.Command{
    Use:   "sha256",
    Short: "Generate SHA256/HmacSHA256 ciphertext",
    Long: `Generate SHA256/HmacSHA256 ciphertext, Generate HmacSHA256 ciphertext with flag(-k/--key)`,
    Args: cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        encodeSHA256(args[0])
    },
}

func encodeSHA256(str string) {
    var result string
    if sha256Key != "" {
        result = util.HmacSHA256(str, sha256Key)
    } else {
        result = util.SHA256(str)
    }
    fmt.Println(result)
    util.WriteClipboard(result)
}

func init() {
    codecCmd.AddCommand(sha256Cmd)
    sha256US = sha256Cmd.UsageString()
    sha256Cmd.SetUsageFunc(sha256UsageFunc)
    sha256Cmd.Flags().StringVarP(&sha256Key, "key", "k", "", "Hmac SHA256 key")
}

func sha256UsageFunc(command *cobra.Command) error {
    fmt.Println(sha256US)
    fmt.Println(`Examples:
  Non-interactive:
    igo codec sha256 oneisx
    igo codec sha256 oneisx -k thanks
    igo codec sha256 oneisx --key thanks
  Interactive:
    codec sha256 oneisx
    codec sha256 oneisx -k thanks
    codec sha256 oneisx --key thanks`)
    return nil
}
