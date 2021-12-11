package cmd

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"igo/util"
	"strings"
)

var number int

var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "generate uuid",
	Long:  `generate uuid`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		generateUUID()
	},
}

func init() {
	rootCmd.AddCommand(uuidCmd)
	uuidCmd.Flags().IntVarP(&number, "num", "n", 1, "The number of UUID that will be generated.")
}

func generateUUID() {
	var uuids []string
	for i := 0; i < number; i++ {
		newUUID, _ := uuid.NewUUID()
		uuids = append(uuids, newUUID.String())
	}
	str := strings.Join(uuids, "\n")
	fmt.Println(str)
	util.WriteClipboard(str)
}
