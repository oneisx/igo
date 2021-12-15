package cmd

import (
    "bufio"
    "fmt"
    "github.com/spf13/cobra"
    "github.com/tidwall/pretty"
    "igo/cst"
    "igo/util"
    "os"
)

var Ugly bool
var jsonUS string

const (
    prettyMode = "pretty"
    uglyMode   = "ugly"
)

var jsonCmd = &cobra.Command{
    Use:   "json",
    Short: "beautify json",
    Long:  `beautify json default, compress json with flag -u(--ugly)
json string should be end with semicolon(;)`,
    Args:  cobra.NoArgs,
    Run: func(cmd *cobra.Command, args []string) {
        HandleJson()
    },
}

func HandleJson() {
    mode := prettyMode
    if Ugly {
        mode = uglyMode
    }
    fmt.Printf("igo>json:%s>", mode)
    input := readJsonFromTerminal()
    doHandleJson(input)
}

func readJsonFromTerminal() string {
    inputReader := bufio.NewReaderSize(os.Stdin, cst.JsonBufSize)
    input, err := inputReader.ReadString(cst.SemicolonDelim)
    if err != nil {
        panic(err)
    }
    return input[:len(input)-1]
}

func doHandleJson(str string) {
    jByte := []byte(str)
    var result []byte
    if Ugly {
        result = pretty.Ugly(jByte)
    } else {
        result = pretty.Pretty(jByte)
    }
    json := string(result)
    fmt.Println(json)
    util.WriteText2Clipboard(json)
}

func SetDefaultPretty() {
    Ugly = false
}

func init() {
    rootCmd.AddCommand(jsonCmd)
    jsonCmd.Flags().BoolVarP(&Ugly, "ugly", "u", false, "compress json")
    jsonUS = jsonCmd.UsageString()
    jsonCmd.SetUsageFunc(jsonCusUsageFunc)
}

func jsonCusUsageFunc(command *cobra.Command) error {
    fmt.Println(jsonUS)
    fmt.Println(`Examples:
  Non-interactive:
    igo json
    igo json -u
  Interactive:
    json
    json -u`)
    return nil
}
