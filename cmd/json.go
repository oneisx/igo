package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tidwall/pretty"
	"igo/constant"
	"igo/util"
	"os"
)

var ugly bool

const (
	prettyMode = "pretty"
	uglyMode   = "ugly"
)

// jsonCmd represents the json command
var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "beautify json",
	Long:  `beautify json`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		handleJson()
	},
}

func handleJson() {
	mode := prettyMode
	if ugly {
		mode = uglyMode
	}
	fmt.Printf("igo>json:%s>", mode)
	input := readJsonFromTerminal()
	doHandleJson(input)
}

func readJsonFromTerminal() string {
	inputReader := bufio.NewReaderSize(os.Stdin, constant.JsonBufSize)
	input, err := inputReader.ReadString(constant.SemicolonDelim)
	if err != nil {
		panic(err)
	}
	return input[:len(input)-1]
}

func doHandleJson(str string) {
	jByte := []byte(str)
	var result []byte
	if ugly {
		result = pretty.Ugly(jByte)
	} else {
		result = pretty.Pretty(jByte)
	}
	json := string(result)
	fmt.Println(json)
	util.WriteClipboard(json)
}

func init() {
	rootCmd.AddCommand(jsonCmd)
	jsonCmd.Flags().BoolVarP(&ugly, "ugly", "u", false, "compress json")
}
