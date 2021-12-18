package cmd

import (
    "github.com/spf13/cobra"
)

// since v1.0.0
var codecCmd = &cobra.Command{
    Use:   "codec",
    Short: "A collection of encryption and decryption tools",
    Long: `A collection of encryption and decryption tools, 
including MD5/HmacMD5, Base64, SHA1/HmacSHA1, SHA256/HmacSHA256, 
SHA512/HmacSHA512, AES, etc.`,
}

func init() {
    rootCmd.AddCommand(codecCmd)
}
