package cmd

import (
    "fmt"
    "igo/util"

    "github.com/spf13/cobra"
)

var sha1US string
var sha1Key string

var sha1Cmd = &cobra.Command{
    Use:   "sha1",
    Short: "Generate SHA1/HmacSHA1 ciphertext",
    Long: `Generate SHA1/HmacSHA1 ciphertext, Generate HmacSHA1 ciphertext with flag(-k/--key)`,
    Args: cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        encodeSHA1(args[0])
    },
}

func encodeSHA1(str string) {
    var result string
    if sha1Key != "" {
        result = util.HmacSHA1(str, sha1Key)
    } else {
        result = util.SHA1(str)
    }
    fmt.Println(result)
    util.WriteText2Clipboard(result)
}

func init() {
    codecCmd.AddCommand(sha1Cmd)
    sha1US = sha1Cmd.UsageString()
    sha1Cmd.SetUsageFunc(sha1UsageFunc)
    sha1Cmd.Flags().StringVarP(&sha1Key, "key", "k", "", "Hmac SHA1 key")
}

func sha1UsageFunc(command *cobra.Command) error {
    fmt.Println(sha1US)
    fmt.Println(`Examples:
  Non-interactive:
    igo codec sha1 oneisx
    igo codec sha1 oneisx -k thanks
    igo codec sha1 oneisx --key thanks
  Interactive:
    codec sha1 oneisx
    codec sha1 oneisx -k thanks
    codec sha1 oneisx --key thanks`)
    return nil
}
