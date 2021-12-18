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
        handleFlags()
    },
}

func Execute() {
    cobra.CheckErr(rootCmd.Execute())
}

func handleFlags() {
    if version {
        printVersion()
    }

    if interactive {
        doInteractive()
    }

    if !version && !interactive {
       util.ExecOSCmd(cst.AppName + cst.SpaceDelim + cst.HelpCommand)
    }
}

func doInteractive() {
    inputReader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("igo>")
        input := readString(inputReader)

        checkExit(input)

        execSubCommand(input)
    }
}

func execSubCommand(input string) {
    if strings.Contains(input, cst.JsonCommand) {
        execJsonCommand(input)
        return
    }

    if input != "" {
        util.ExecOSCmd(cst.AppName + cst.SpaceDelim + input)
    }
}

func execJsonCommand(input string) {
    defer SetDefaultPretty()
    ctx := strings.Split(input, cst.SpaceDelim)
    for _, v := range ctx {
        if v == cst.UglyFlag {
            Ugly = true
        }
        if v != "" && v != cst.JsonCommand && v != cst.UglyFlag {
            util.ExecOSCmd(cst.AppName + cst.SpaceDelim + input)
            return
        }
    }
    HandleJson()
}

func readString(reader *bufio.Reader) string {
    input, err := reader.ReadString('\n')
    if err != nil {
        os.Exit(1)
    }
    return util.RemoveLineBreak(input)
}

func checkExit(input string) {
    if strings.Contains(input, cst.QFlag) || strings.Contains(input, cst.QuitFlag) {
        fmt.Println("bye")
        os.Exit(1)
    }
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
