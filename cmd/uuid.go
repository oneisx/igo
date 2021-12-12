package cmd

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"github.com/spf13/cobra"
	"igo/util"
	"strings"
)

var number int
var uuidUS string

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
    uuidUS = uuidCmd.UsageString()
	uuidCmd.SetUsageFunc(uuidUsageFunc)
}

func uuidUsageFunc(cmd *cobra.Command) error {
	fmt.Println(uuidUS)
	fmt.Println(`Examples:
  Non-interactive:
    igo uuid
    igo uuid -n 3
    igo uuid --num 3
  Interactive:
    uuid
    uuid -n 3
    uuid --num 3`)
	return nil
}

func generateUUID() {
	var uuids []string
	for i := 0; i < number; i++ {
		newUUID := UUID()
		uuids = append(uuids, newUUID)
	}
	str := strings.Join(uuids, "\n")
	fmt.Println(str)
	util.WriteClipboard(str)
}

func UUID() string {
	var first, last uint32
	var middle [4]uint16
	randomBytes := RandomBytes(16)
	buffer := bytes.NewBuffer(randomBytes)
	_ = binary.Read(buffer, binary.BigEndian, &first)
	for i := 0; i < 4; i++ {
		_ = binary.Read(buffer, binary.BigEndian, &middle[i])
	}
	_ = binary.Read(buffer, binary.BigEndian, &last)
	middle[1] = (middle[1] & 0x0fff) | 0x4000
	middle[2] = (middle[2] & 0x3fff) | 0x8000
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%04x%08x",
		first, middle[0], middle[1], middle[2], middle[3], last)
}

func RandomBytes(n int) []byte {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return b
}
