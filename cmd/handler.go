package cmd

import (
    "fmt"
    "igo/cst"
    "igo/util"
    "os"
    "strings"
)

var HandlerList = []Handler{new(BaseHandler), new(IgoHandler)}
var BaseCmdMap = map[string]BaseCmd{"cd": new(ChangeDirectory), "pwd": new(PrintWorkDirectory), "ll": new(ListDirectory)}

type Handler interface {
    Handle(command string) bool
}

type BaseHandler struct{}

type IgoHandler struct{}

func (baseHandler *BaseHandler) Handle(command string) bool {
    arr := strings.Split(command, cst.SpaceDelim)
    if c, ok := BaseCmdMap[arr[0]]; ok {
        c.exec(arr[0], arr[1:])
        return true
    }
    return false
}

func (igoHandler *IgoHandler) Handle(command string) bool {
    execIgoSubCommand(command)
    return false
}

func execIgoSubCommand(input string) {
    checkExit(input)
    
    if strings.Contains(input, cst.JsonCommand) {
        ExecJsonCommand(input)
        return
    }
    
    if strings.Contains(input, cst.SqlCommand) {
        ExecSqlCommand(input)
        return
    }
    
    if input != "" {
        util.Operator.ExecOSCmd(cst.AppName + cst.SpaceDelim + input)
    }
}

func checkExit(input string) {
    if strings.Contains(input, cst.QFlag) || strings.Contains(input, cst.QuitFlag) {
        fmt.Println("bye")
        os.Exit(1)
    }
}
