package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "igo",
	Short: "igo",
	Long:  `igo`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	interactiveModeIfNeed(rootCmd)
	cobra.CheckErr(rootCmd.Execute())
}

func interactiveModeIfNeed(cmd *cobra.Command) {
	isInteractive, _ := cmd.Flags().GetBool("interactive")
	if isInteractive {
		doInteractive(cmd)
	}
}

func doInteractive(cmd *cobra.Command) {
	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("igo>")
		input, err := inputReader.ReadString('\n')
		checkExit(input)
		if err != nil {
			return
		}
		fmt.Print(input)
	}
}

func checkExit(input string) {
	if strings.Compare(":q\n", input) == 0 {
		fmt.Println("bye")
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.igo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolP("interactive", "i", false, "Interactively execute commands")
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
