package cmd

import (
    "bufio"
    "fmt"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "igo/cst"
    "igo/util"
    "os"
    "runtime"
    "strings"
)

var cfgFile string
var interactive bool
var version bool
var quit bool

// since v1.0.0

var rootCmd = &cobra.Command{
    Use:   "igo",
    Short: "igo",
    Long: `Welcome to igo, igo is a good helper, can generate UUID, beautify JSON, convert timestamp, etc.
No need to copy manually, the generated content will be placed in the clipboard.
For more functions, please see the help.`,
    Run: func(cmd *cobra.Command, args []string) {
        execCmd()
    },
}

func Execute() {
    cobra.CheckErr(rootCmd.Execute())
}

func execCmd() {
    if version {
        printVersion()
    }

    if interactive {
        doInteractive()
    }

    if !version && !interactive {
       util.Operator.ExecOSCmd(cst.AppName + cst.SpaceDelim + cst.HelpCommand)
    }
}

func doInteractive() {
    printVersion()
    inputReader := bufio.NewReader(os.Stdin)
    for {
        fmt.Printf("igo:%s>", util.Operator.ShortPath())
        //fmt.Print("igo>")
        input := readString(inputReader)
    
        for _, h := range HandlerList {
            if h.Handle(input) {
                break
            }
        }
    }
}

func readString(reader *bufio.Reader) string {
    line, _, _ := reader.ReadLine()
    return strings.TrimSpace(string(line))
}

func init() {
    cobra.OnInitialize(initConfig)
    rootCmd.CompletionOptions.DisableDefaultCmd = true
    rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.igo.yaml)")
    rootCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "interactively execute commands")
    rootCmd.Flags().BoolVarP(&version, "version", "v", false, "print the version of igo")
    rootCmd.Flags().BoolVarP(&quit, "quit", "q", false, "quit interactive mode")
}

func initConfig() {
    if cfgFile != "" {
        // Use config file from the flag.
        viper.SetConfigFile(cfgFile)
    } else {
        // Find home directory.
        home, err := os.UserHomeDir()
        cobra.CheckErr(err)

        // Search config in home directory with name ".igo" (without extension).
        viper.AddConfigPath(home)
        viper.SetConfigType("yaml")
        viper.SetConfigName(".igo")
    }

    viper.AutomaticEnv() // read in environment variables that match

    // If a config file is found, read it in.
    if err := viper.ReadInConfig(); err == nil {
        _, err := fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
        if err != nil {
            return
        }
    }
}

func printVersion() {
    fmt.Println(cst.AppName + cst.SpaceDelim + cst.AppVersion + cst.SpaceDelim + runtime.GOOS + cst.SlashDelim + runtime.GOARCH + cst.CommaDelim + cst.PoweredBy)
}

