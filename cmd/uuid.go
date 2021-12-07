package cmd

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"igo/util"
	"strings"
)

// uuidCmd represents the uuid command
var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generate UUID",
	Long:  `Generate UUID`,
	Run: func(cmd *cobra.Command, args []string) {
		generateUUID(args)
	},
}

func init() {
	rootCmd.AddCommand(uuidCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uuidCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uuidCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func generateUUID(args []string) {
	length := len(args)
	if length > 1 {
		err := fmt.Errorf("too many args found, length: %d", length)
		fmt.Println(err.Error())
		return
	}

	num := 0
	if length == 1 && util.IsNum(args[0]) {
		num = util.Atoi(args[0])
	}

	if length == 0 {
		num = 1
	}
	doGenerateUUID(num)
}

func doGenerateUUID(num int) {
	var uuids []string
	for i := 0; i < num; i++ {
		newUUID, _ := uuid.NewUUID()
		uuids = append(uuids, newUUID.String())
	}
	str := strings.Join(uuids, "\n")
	fmt.Println(str)
	util.WriteClipboard(str)
}

