package cmd

import (
    "bufio"
    "fmt"
    "github.com/spf13/cobra"
    "igo/cst"
    "igo/util"
    "os"
    "strconv"
    "strings"
)

var (
    list        bool
    searchKey   string
    sqlKey      string
    updateSqlId int
    delSqlId    int
    pickSqlId   int
)

var sqlCmd = &cobra.Command{
    Use:   "sql",
    Short: "Memo function designed for SQL",
    Long:  `Memo function designed for SQL`,
    Run: func(cmd *cobra.Command, args []string) {
        execOperation()
    },
}

func init() {
    rootCmd.AddCommand(sqlCmd)
    sqlCmd.Flags().BoolVarP(&list, "list", "l", false, "list memo data")
    sqlCmd.Flags().StringVarP(&searchKey, "search", "s", "", "search memo data")
    sqlCmd.Flags().StringVarP(&sqlKey, "add", "a", "", "add memo data")
    sqlCmd.Flags().IntVarP(&updateSqlId, "update", "u", -1, "update memo data")
    sqlCmd.Flags().IntVarP(&delSqlId, "del", "d", -1, "del memo data")
    sqlCmd.Flags().IntVarP(&pickSqlId, "pick", "p", -1, "select memo data")
}

func ExecSqlCommand(input string) {
    resolveInput(input)
    execOperation()
    resetField()
}

func resetField() {
    list=false
    searchKey=""
    sqlKey=""
    updateSqlId=-1
    delSqlId=-1
    pickSqlId=-1
}

func resolveInput(input string) {
    param := strings.Split(input, cst.SpaceDelim)[1:]
    switch param[0] {
    case "-l":
        fallthrough
    case "--list":
        list = true
    case "-s":
        fallthrough
    case "--search":
        searchKey = param[1]
    case "-a":
        fallthrough
    case "--add":
        sqlKey = param[1]
    case "-u":
        fallthrough
    case "--update":
        updateSqlId, _ = strconv.Atoi(param[1])
    case "-d":
        fallthrough
    case "--del":
        delSqlId, _ = strconv.Atoi(param[1])
    case "-p":
        fallthrough
    case "--pick":
        pickSqlId, _ = strconv.Atoi(param[1])
    }
}

func execOperation() {
    o := chooseOperation()
    o.exec()
}

func chooseOperation() operation {
    if list {
        return new(listOperation)
    }
    if searchKey != "" {
        return new(searchOperation)
    }
    if pickSqlId != -1 {
        return new(pickOperation)
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
    return new(helpOperation)
}

type operation interface {
    exec()
}

type listOperation struct{}
type searchOperation struct{}
type addOperation struct{}
type updateOperation struct{}
type delOperation struct{}
type pickOperation struct{}
type helpOperation struct{}

func (l *listOperation) exec() {
    listSql("")
}

func (s *searchOperation) exec() {
    listSql(searchKey)
}

func (a *addOperation) exec() {
    fmt.Printf("igo>sql:add:%s>", sqlKey)
    sql := readSqlFromTerminal()
    util.PutSQL(util.MemoData{Key: sqlKey, Data: sql})
    fmt.Println("sql saved successfully!")
}

func (u *updateOperation) exec() {
    m := util.GetSQL(updateSqlId)
    if m.Data == "" {
        err := fmt.Errorf("error: sql can not be found with id=%v, failed to update", updateSqlId)
        fmt.Println(err)
        return
    }
    fmt.Printf("igo>sql:update:%s>", m.Key)
    sql := readSqlFromTerminal()
    m.Data = sql
    util.UpdateSQL(m, updateSqlId)
    fmt.Println("sql update successfully!")
}

func (d *delOperation) exec() {
    util.DelSQL(delSqlId)
    fmt.Println("delete sql successfully!")
}

func (p *pickOperation) exec() {
    m := util.GetSQL(pickSqlId)
    if m.Data == "" {
        err := fmt.Errorf("error: sql can not be found with id=%v", pickSqlId)
        fmt.Println(err)
        return
    }
    fmt.Println(m.Data)
    util.WriteText2Clipboard(m.Data)
}

func (h *helpOperation) exec() {
    util.ExecOSCmd(cst.AppName + cst.SpaceDelim + cst.SqlCommand + cst.SpaceDelim + cst.HelpFlag)
}

func readCommandFromTerminal4List() string {
    inputReader := bufio.NewReaderSize(os.Stdin, cst.SqlBufSize)
    input, err := inputReader.ReadString('\n')
    if err != nil {
        panic(err)
    }
    return input[:len(input)-1]
}

func readSqlFromTerminal() string {
    inputReader := bufio.NewReaderSize(os.Stdin, cst.SqlBufSize)
    input, err := inputReader.ReadString(cst.SemicolonDelim)
    if err != nil {
        panic(err)
    }
    return input[:len(input)-1]
}

func listSql(key string) {
    ms := util.GetAllSQL()
    mdSlice := convertAndFilterSql(ms, key)
    fmt.Println("(", len(mdSlice), "rows)")
    count := 0
    page := 1
    for _, memoData := range mdSlice {
        if count == 0 {
            fmt.Println("page:", page)
        }
        fmt.Println("id:", memoData.Id, cst.SpaceDelim, "key:", memoData.Key)
        count++
        if count >= 10 {
            fmt.Println("Press 'Enter' to continue and 'q' to end browsing")
            c := readCommandFromTerminal4List()
            if c == "q" {
                break
            }
            count = 0
            page++
        }
    }
}

func convertAndFilterSql(ms map[int]util.MemoData, key string) []util.MemoData {
    var mdSlice []util.MemoData
    for i := 1; i <= len(ms); i++ {
        m := ms[i]
        if key == "" {
            mdSlice = append(mdSlice, m)
        } else if key != "" && strings.Contains(m.Key, key) {
            mdSlice = append(mdSlice, m)
        }
    }
    return mdSlice
}
