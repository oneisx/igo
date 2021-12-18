package cmd

import (
    "fmt"
    "igo/util"

    "github.com/spf13/cobra"
)

var md5US string
var md5Key string

// since v1.0.0
var md5Cmd = &cobra.Command{
    Use:   "md5",
    Short: "Generate MD5/HmacMD5 ciphertext",
    Long: `Generate MD5/HmacMD5 ciphertext, Generate HmacMD5 ciphertext with flag(-k/--key)`,
    Args: cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        encodeMD5(args[0])
    },
}

func encodeMD5(str string) {
    var result string
    if md5Key != "" {
        result = util.HmacMD5(str, md5Key)
    } else {
        result = util.MD5(str)
    }
    fmt.Println(result)
    util.WriteText2Clipboard(result)
}

func init() {
    codecCmd.AddCommand(md5Cmd)
    md5US = md5Cmd.UsageString()
    md5Cmd.SetUsageFunc(md5UsageFunc)
    md5Cmd.Flags().StringVarP(&md5Key, "key", "k", "", "Hmac MD5 key")
}

func md5UsageFunc(command *cobra.Command) error {
    fmt.Println(md5US)
    fmt.Println(`Examples:
  Non-interactive:
    igo codec md5 oneisx
    igo codec md5 oneisx -k thanks
    igo codec md5 oneisx --key thanks
  Interactive:
    codec md5 oneisx
    codec md5 oneisx -k thanks
    codec md5 oneisx --key thanks`)
    return nil
}
