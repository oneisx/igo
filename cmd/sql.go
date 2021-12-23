package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"igo/cst"
	"igo/util"
	"os"
)

var (
	list bool
	search string
	sqlKey string
	updateSqlId int
	delSqlId int
	pickSqlId int
	//quit bool
)

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "Memo function designed for SQL",
	Long: `Memo function designed for SQL`,
	Run: func(cmd *cobra.Command, args []string) {
		o := chooseOperation()
		o.exec()
	},
}

func init() {
	rootCmd.AddCommand(sqlCmd)
	sqlCmd.Flags().BoolVarP(&list, "list", "l", false, "list memo data")
	sqlCmd.Flags().StringVarP(&search, "search", "s", "", "search memo data")
	sqlCmd.Flags().StringVarP(&sqlKey, "add", "a", "", "add memo data")
	sqlCmd.Flags().IntVarP(&updateSqlId, "update", "u", -1, "update memo data")
	sqlCmd.Flags().IntVarP(&delSqlId, "del", "d", -1, "del memo data")
	sqlCmd.Flags().IntVarP(&pickSqlId, "pick", "p", -1, "select memo data")
	//sqlCmd.Flags().BoolVarP(&quit, "quit", "q", false, "quit sql command")
}

func chooseOperation() operation {
	if list {
		return new(listOperation)
	}
	if search != "" {
		return new(searchOperation)
	}
	if sqlKey != "" {
		return new(addOperation)
	}
	if updateSqlId != -1 {
		return new(updateOperation)
	}
	if delSqlId != -1 {
		return new(delOperation)
	}
	if pickSqlId != -1 {
		return new(pickOperation)
	}
	return new(helpOperation)
}

type operation interface{
	exec()
}

type listOperation struct {}
type searchOperation struct {}
type addOperation struct {}
type updateOperation struct {}
type delOperation struct {}
type pickOperation struct {}
type helpOperation struct {}

func (l *listOperation) exec() {

}

func (s *searchOperation) exec() {

}

func (a *addOperation) exec() {
	fmt.Printf("igo>sql:add:%s>", sqlKey)
	sql := readSqlFromTerminal()
	util.PutSQL(util.MemoData{Key: sqlKey, Data: sql})
	fmt.Println("SQL saved successfully!")
}

func (u *updateOperation) exec() {

}

func (d *delOperation) exec() {

}

func (p *pickOperation) exec() {

}

func (h *helpOperation) exec() {
	util.ExecOSCmd(cst.AppName + cst.SpaceDelim + cst.SqlCommand + cst.SpaceDelim + cst.HelpFlag)
}

func readSqlFromTerminal() string {
	inputReader := bufio.NewReaderSize(os.Stdin, cst.SqlBufSize)
	input, err := inputReader.ReadString(cst.SemicolonDelim)
	if err != nil {
		panic(err)
	}
	return input[:len(input)-1]
}