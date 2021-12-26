package cmd

import (
    "bufio"
    "fmt"
    "github.com/spf13/cobra"
    "igo/cst"
    "igo/util"
    "os"
    "sort"
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
    a := strings.Split(input, cst.SpaceDelim)[1:]
    a = removeEmptyStr(a)
    _ = sqlCmd.ParseFlags(a)
    execOperation()
    resetField()
}

func resetField() {
    list = false
    searchKey = ""
    sqlKey = ""
    updateSqlId = -1
    delSqlId = -1
    pickSqlId = -1
}

func removeEmptyStr(a []string) []string {
    var newA []string
    for _, v := range a {
        if v != "" {
            newA = append(newA, v)
        }
    }
    return newA
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
    listSql("", "list")
}

func (s *searchOperation) exec() {
    listSql(searchKey,"search")
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
    fmt.Println("sql to be updated:")
    fmt.Println(m.Data)
    fmt.Printf("igo>sql:update:%s>", m.Key)
    sql := readSqlFromTerminal()
    m.Data = sql
    util.UpdateSQL(m, updateSqlId)
    fmt.Println("update sql successfully!")
}

func (d *delOperation) exec() {
    util.DelSQL(delSqlId)
    fmt.Println("delete sql successfully!")
}

func (p *pickOperation) exec() {
    pickSQL(pickSqlId)
}

func (h *helpOperation) exec() {
    util.ExecOSCmd(cst.AppName + cst.SpaceDelim + cst.SqlCommand + cst.SpaceDelim + cst.HelpFlag)
}

func pickSQL(id int) {
    m := util.GetSQL(id)
    if m.Data == "" {
        err := fmt.Errorf("error: sql can not be found with id=%v", pickSqlId)
        fmt.Println(err)
        return
    }
    fmt.Println(m.Data)
    util.WriteText2Clipboard(m.Data)
}

func readCommandFromTerminal4List(flag string) string {
    fmt.Printf("igo>sql:%s>", flag)
    inputReader := bufio.NewReaderSize(os.Stdin, cst.SqlBufSize)
    line, _, _ := inputReader.ReadLine()
    return string(line)
}

func readSqlFromTerminal() string {
    inputReader := bufio.NewReaderSize(os.Stdin, cst.SqlBufSize)
    input, err := inputReader.ReadString(cst.SemicolonDelim)
    if err != nil {
       panic(err)
    }
    return input[:len(input)-1]
}

func listSql(key string, flag string) {
    ms := util.GetAllSQL()
    mdSlice := convertAndFilterSql(ms, key)
    fmt.Println("(", len(mdSlice), "rows )")
    count := 0
    page := 1
    for _, memoData := range mdSlice {
        if count == 0 {
            fmt.Println("page:", page)
        }
        fmt.Println("id:", memoData.Id, cst.SpaceDelim, "key:", memoData.Key)
        count++
        if count >= 10 {
            fmt.Println("(Pick: <id> / PaDn: Enter / Quit: q)")
            c := readCommandFromTerminal4List(flag)
            if strings.Compare(c, "q") == 0 {
                break
            }
            if util.IsNum(c) {
                pickSQL(util.ParseInt(c))
                return
            }
            count = 0
            page++
        }
    }
    fmt.Println("(Pick: <id> / Quit: Enter)")
    c := readCommandFromTerminal4List(flag)
    if util.IsNum(c) {
        pickSQL(util.ParseInt(c))
    }
}

func convertAndFilterSql(ms map[int]util.MemoData, key string) []util.MemoData {
    var mdSlice MemoDataSlice
    for _, memoData := range ms {
        if key == "" {
            mdSlice = append(mdSlice, memoData)
        } else if key != "" && strings.Contains(memoData.Key, key) {
            mdSlice = append(mdSlice, memoData)
        }
    }
    sort.Sort(mdSlice)
    return mdSlice
}

type MemoDataSlice []util.MemoData

func (s MemoDataSlice) Len() int           { return len(s) }
func (s MemoDataSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s MemoDataSlice) Less(i, j int) bool { return s[i].Id < s[j].Id }
