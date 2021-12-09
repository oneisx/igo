package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"os"
	"os/exec"
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
		input := readString(inputReader)

		checkExit(input)

		execChildCommand(cmd, input)
	}
}

func execChildCommand(cmd *cobra.Command, input string) {
	command := "cmd.exe"
	params := []string{"/c", "igo " + input}
	execCommand(command, params)
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

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.igo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolP("interactive", "i", true, "Interactively execute commands")
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

func execCommand(commandName string, params []string) bool {
	cmd := exec.Command(commandName, params...)

	//显示运行的命令
	//fmt.Println(cmd.Args)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		return false
	}

	err = cmd.Start()
	if err != nil {
		return false
	}

	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}

	err = cmd.Wait()
	if err != nil {
		return false
	}
	return true
}
