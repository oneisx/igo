package cmd

import (
    "fmt"
    "igo/util"

    "github.com/spf13/cobra"
)

var decryptAES bool
var aesUS string
var aesKey string

var aesCmd = &cobra.Command{
    Use:   "aes",
    Short: "Generate aes ciphertext、decrypt aes ciphertext",
    Long:  `Generate aes ciphertext、decrypt aes ciphertext, use flag (-d/--decrypt) to decrypt`,
    Args:  cobra.ExactArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
        return doAesCodec(args[0])
    },
}

func doAesCodec(str string) error {
    var result string
    var err error
    if decryptAES {
        result, err = util.DecryptAES(str, aesKey)
    } else {
        result, err = util.EncryptAES(str, aesKey)
    }
    if err != nil {
        return err
    }
    fmt.Println(result)
    util.WriteClipboard(result)
    return nil
}

func init() {
    codecCmd.AddCommand(aesCmd)
    aesCmd.Flags().BoolVarP(&decryptAES, "decrypt", "d", false, "decrypt aes ciphertext")
    aesCmd.Flags().StringVarP(&aesKey, "key", "k", "0000000000000000", "the key of AES, " +
        "length should be in (16,24,32), the default value is not recommended.")
    aesUS = aesCmd.UsageString()
    aesCmd.SetUsageFunc(aesUsageFunc)
}

func aesUsageFunc(command *cobra.Command) error {
    fmt.Println(aesUS)
    fmt.Println(`Examples:
  Non-interactive:
    igo codec aes oneisx
    igo codec aes uZERhkcVewZ7S1j1co+QSkKdvf/52DqkDXgAcJktido= -d
    igo codec aes oneisx -k 52DqkDXgAcJktido (Recommend custom key)
    igo codec aes OYA/OY1bj6J1wRywYYCIwMC9oW8RqoByngxsBUlGhuw= -dk 52DqkDXgAcJktido (Recommend custom key)
  Interactive:
    codec aes oneisx
    codec aes -d uZERhkcVewZ7S1j1co+QSkKdvf/52DqkDXgAcJktido= -d
    codec aes oneisx -k 52DqkDXgAcJktido (Recommend custom key)
    codec aes OYA/OY1bj6J1wRywYYCIwMC9oW8RqoByngxsBUlGhuw= -dk 52DqkDXgAcJktido (Recommend custom key)`)
    return nil
}
