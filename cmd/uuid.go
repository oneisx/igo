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
	Short: "generate uuid",
	Long:  `generate uuid`,
	Run: func(cmd *cobra.Command, args []string) {
		generateUUID(args)
	},
}

func init() {
	rootCmd.AddCommand(uuidCmd)
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

