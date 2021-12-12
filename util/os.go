package util

import (
    "os"
    "os/exec"
)

const (
    winOSCommand    = "cmd.exe"
    winClearCommand = "cls"
)

func ClearScreen() {
    ExecOSCmd(winClearCommand)
}

func ExecOSCmd(param string) bool {
    params := []string{"/c", param}
    cmd := exec.Command(winOSCommand, params...)

    //显示运行的命令
    //fmt.Println(cmd.Args)

    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
        return false
    }
    return true
}
