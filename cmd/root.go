package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"igo/util"
	"os"
	"strings"
)

var cfgFile string
var interactive bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "igo",
	Short: "igo",
	Long:  `igo`,
	Run: func(cmd *cobra.Command, args []string) {
		interactiveModeIfNeed()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func interactiveModeIfNeed() {
	if interactive {
		doInteractive()
	}
}

func doInteractive() {
	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("igo>")
		input := readString(inputReader)

		checkExit(input)

		if input == "json" {
			fmt.Println("json command does not support interactive mode")
			continue
		}

		execChildCommand(input)
	}
}

func execChildCommand(input string) {
	if input != "" {
		util.ExecOSCmd("igo " + input)
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.igo.yaml)")
	rootCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "interactively execute commands")
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
