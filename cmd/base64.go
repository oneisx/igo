package cmd

import (
    "fmt"
    "igo/util"

    "github.com/spf13/cobra"
)

var decryptBase64 bool
var base64US string

var base64Cmd = &cobra.Command{
    Use:   "base64",
    Short: "Generate base64 ciphertext、decrypt base64 ciphertext",
    Long:  `Generate base64 ciphertext、decrypt base64 ciphertext, use flag (-d/--decrypt) to decrypt`,
    Args:  cobra.ExactArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
        return doBase64Codec(args[0])
    },
}

func doBase64Codec(str string) error {
    var result string
    var err error
    if decryptBase64 {
        result, err = util.DecryptBase64(str)
        if err != nil {
            return fmt.Errorf("invalid base64 ciphertext")
        }
    } else {
        result = util.EncryptBase64(str)
    }
    fmt.Println(result)
    util.WriteClipboard(result)
    return nil
}

func init() {
    codecCmd.AddCommand(base64Cmd)
    base64Cmd.Flags().BoolVarP(&decryptBase64, "decrypt", "d", false, "decrypt base64 ciphertext")
    base64US = base64Cmd.UsageString()
    base64Cmd.SetUsageFunc(base64UsageFunc)
}

func base64UsageFunc(command *cobra.Command) error {
    fmt.Println(base64US)
    fmt.Println(`Examples:
  Non-interactive:
    igo codec base64 oneisx
    igo codec base64 -d b25laXN4
  Interactive:
    codec base64 oneisx
    codec base64 -d b25laXN4`)
    return nil
}
