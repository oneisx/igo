package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"igo/constant"
	"igo/util"
	"os"
	"runtime"
	"strings"
)

var cfgFile string
var interactive bool
var version bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "igo",
	Short: "igo",
	Long:  `igo`,
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
		util.ExecOSCmd(constant.IgoHelpCommand)
	}
}

func doInteractive() {
	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("igo>")
		input := readString(inputReader)

		checkExit(input)

		if input == "json" {
			fmt.Println("Error: json command does not support interactive mode")
			continue
		}

		execChildCommand(input)
	}
}

func execChildCommand(input string) {
	if input != "" {
		util.ExecOSCmd(constant.AppName + constant.SpaceDelim + input)
	}
}

func readString(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')
	if err != nil {
		os.Exit(1)
	}
	return input[:len(input)-2]
}

func checkExit(input string) {
	if strings.Compare(":q", input) == 0 {
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
}

// initConfig reads in config file and ENV variables if set.
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
	fmt.Println(constant.AppName + constant.SpaceDelim + constant.AppVersion + constant.SpaceDelim + runtime.GOOS + constant.SlashDelim + runtime.GOARCH + constant.CommaDelim + constant.PoweredBy)
}