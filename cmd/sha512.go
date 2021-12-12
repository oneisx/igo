package cmd

import (
	"fmt"
	"igo/util"

	"github.com/spf13/cobra"
)

var sha512US string
var sha512Key string

var sha512Cmd = &cobra.Command{
	Use:   "sha512",
	Short: "Generate SHA512/HmacSHA512 ciphertext",
	Long: `Generate SHA512/HmacSHA512 ciphertext, Generate HmacSHA512 ciphertext with flag(-k/--key)`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		encodeSHA512(args[0])
	},
}

func encodeSHA512(str string) {
	var result string
	if sha512Key != "" {
		result = util.HmacSHA512(str, sha512Key)
	} else {
		result = util.SHA512(str)
	}
	fmt.Println(result)
	util.WriteClipboard(result)
}

func init() {
	codecCmd.AddCommand(sha512Cmd)
	sha512US = sha512Cmd.UsageString()
	sha512Cmd.SetUsageFunc(sha512UsageFunc)
	sha512Cmd.Flags().StringVarP(&sha512Key, "key", "k", "", "Hmac SHA512 key")
}

func sha512UsageFunc(command *cobra.Command) error {
	fmt.Println(sha512US)
	fmt.Println(`Examples:
  Non-interactive:
    igo codec sha512 oneisx
	igo codec sha512 oneisx -k thanks
	igo codec sha512 oneisx --key thanks
  Interactive:
    codec sha512 oneisx
	codec sha512 oneisx -k thanks
	codec sha512 oneisx --key thanks`)
	return nil
}
