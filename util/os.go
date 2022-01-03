package util

import (
    "fmt"
    "os"
    "os/exec"
    "runtime"
    "strings"
)

const (
    windows          = "windows"
    winOSCommand     = "cmd.exe"
    winCommandOption = "/c"
    winClearCommand  = "cls"
    
    linux              = "linux"
    linuxOSCommand     = "/bin/bash"
    linuxCommandOption = "-c"
    
    macos              = "darwin"
    macOSCommand       = "/usr/bin/open"
    macOSCommandOption = "-a"
    
    ClearCommand = "clear"
    CdCommand    = "cd"
    
    userHomeDirSymbol = "~"
)

var Operator = initOS()

type OperationSystem interface {
    ClearScreen()
    ExecOSCmd(command string)
    GetLineFeed() string
    RmLineFeed(str string) string
    buildCmd(command string) *exec.Cmd
    ShortPath() string
    PrintWorkDirectory() string
    ChangeDirectory(path string)
    ListDirectory()
    init() OperationSystem
}

type Windows struct {
    CurrentWorkDirectory string
    LastWorkDirectory    string
}

type Linux struct {
    CurrentWorkDirectory string
    LastWorkDirectory    string
}

type MacOS struct {
    CurrentWorkDirectory string
    LastWorkDirectory    string
}

func (windows *Windows) ClearScreen() {
    cmd := windows.buildCmd(winClearCommand)
    doExecOSCmd(cmd)
}

func (linux *Linux) ClearScreen() {
    cmd := linux.buildCmd(ClearCommand)
    doExecOSCmd(cmd)
}

func (macos *MacOS) ClearScreen() {
    cmd := macos.buildCmd(ClearCommand)
    doExecOSCmd(cmd)
}

func (windows *Windows) ExecOSCmd(command string) {
    cmd := windows.buildCmd(command)
    ret := doExecOSCmd(cmd)
    if !ret {
        fmt.Println("error: doExecOsCmd failed!")
    }
}

func (linux *Linux) ExecOSCmd(command string) {
    cmd := linux.buildCmd(command)
    ret := doExecOSCmd(cmd)
    if !ret {
        fmt.Println("error: doExecOsCmd failed!")
    }
}

func (macos *MacOS) ExecOSCmd(command string) {
    cmd := macos.buildCmd(command)
    ret := doExecOSCmd(cmd)
    if !ret {
        fmt.Println("error: doExecOsCmd failed!")
    }
}

func (windows *Windows) GetLineFeed() string {
    return "\r\n"
}

func (linux *Linux) GetLineFeed() string {
    return "\n"
}

func (macos *MacOS) GetLineFeed() string {
    return "\r"
}

func (windows *Windows) RmLineFeed(str string) string {
    return str[:len(str)-2]
}

func (linux *Linux) RmLineFeed(str string) string {
    return str[:len(str)-1]
}

func (macos *MacOS) RmLineFeed(str string) string {
    return str[:len(str)-1]
}

func (windows *Windows) PrintWorkDirectory() string {
    pwd := windows.CurrentWorkDirectory
    pwd = "/" + strings.Replace(pwd, ":\\", "/", 1)
    pwd = strings.ReplaceAll(pwd, "\\", "/")
    return pwd
}

func (linux *Linux) PrintWorkDirectory() string {
    return linux.CurrentWorkDirectory
}

func (macos *MacOS) PrintWorkDirectory() string {
    return macos.CurrentWorkDirectory
}

func (windows *Windows) ShortPath() string {
    pwd := windows.PrintWorkDirectory()
    //if strings.Compare(pwd, UserHomeDir()) == 0 {
    //    return userHomeDirSymbol
    //}
    return pwd[strings.LastIndex(pwd, "/")+1:]
}

func (linux *Linux) ShortPath() string {
    pwd := linux.PrintWorkDirectory()
    if strings.Compare(pwd, UserHomeDir()) == 0 {
        return userHomeDirSymbol
    }
    return pwd[strings.LastIndex(pwd, string(os.PathSeparator))+1:]
}

func (macos *MacOS) ShortPath() string {
    pwd := macos.PrintWorkDirectory()
    if strings.Compare(pwd, UserHomeDir()) == 0 {
        return userHomeDirSymbol
    }
    return pwd[strings.LastIndex(pwd, string(os.PathSeparator))+1:]
}

func (windows *Windows) ChangeDirectory(path string) {
    if strings.Index(path, "\\") != -1 {
        fmt.Println("error: invalid path!")
        return
    }
    var tmpPath string
    if path == "-" {
        tmp := windows.CurrentWorkDirectory
        windows.CurrentWorkDirectory = windows.LastWorkDirectory
        windows.LastWorkDirectory = tmp
        return
    }
    if path == "~" {
        windows.LastWorkDirectory = windows.CurrentWorkDirectory
        windows.CurrentWorkDirectory = UserHomeDir()
        return
    }
    if strings.HasPrefix(path, "/") {
        path = strings.Replace(path[1:], "/", ":\\", 1)
        path = strings.ReplaceAll(path, "/", "\\")
        tmpPath = path
    } else {
        tmpPath = windows.CurrentWorkDirectory
        arr := strings.Split(path, "/")
        for _, v := range arr {
            if v == ".." {
                tmpPath = tmpPath[:strings.LastIndex(tmpPath, "\\")]
                if strings.Index(tmpPath, "\\") == -1 {
                    tmpPath = tmpPath + "\\"
                }
                continue
            }
            if v != "" {
                tmpPath = tmpPath + "\\" + v
            }
        }
    }
    if !PathExists(tmpPath) {
        fmt.Println("error: path not exists!")
        return
    }
    if IsFile(tmpPath) {
        fmt.Println("error: path is a File!")
        return
    }
    windows.CurrentWorkDirectory = tmpPath
}

func (linux *Linux) ChangeDirectory(path string) {
    linux.ExecOSCmd("cd " + path)
}

func (macos *MacOS) ChangeDirectory(path string) {
    macos.ExecOSCmd("cd " + path)
}

func (windows *Windows) ListDirectory() {
    pwd := windows.CurrentWorkDirectory
    disk := pwd[:strings.Index(pwd, ":")+1]
    cmd := windows.buildCmd(disk + " && cd " + windows.CurrentWorkDirectory + " && dir")
    ret := doExecOSCmd(cmd)
    if !ret {
        fmt.Println("error: doExecOsCmd failed!")
    }
}

func (linux *Linux) ListDirectory() {
    //linux.ExecOSCmd(CdCommand + cst.SpaceDelim + path)
}

func (macos *MacOS) ListDirectory() {
    //macos.ExecOSCmd(CdCommand + cst.SpaceDelim + path)
}

func (windows *Windows) init() OperationSystem {
    windows.CurrentWorkDirectory = UserHomeDir()
    windows.LastWorkDirectory = UserHomeDir()
    return windows
}

func (linux *Linux) init() OperationSystem {
    linux.CurrentWorkDirectory = UserHomeDir()
    linux.LastWorkDirectory = UserHomeDir()
    return linux
}

func (macos *MacOS) init() OperationSystem {
    macos.CurrentWorkDirectory = UserHomeDir()
    macos.LastWorkDirectory = UserHomeDir()
    return macos
}

func initOS() OperationSystem {
    switch runtime.GOOS {
    case windows:
        return new(Windows).init()
    case linux:
        return new(Linux).init()
    case macos:
        return new(MacOS).init()
    default:
        fmt.Println("Error: Operation system is not supported!")
        os.Exit(1)
    }
    return nil
}

func (macos *MacOS) buildCmd(command string) *exec.Cmd {
    commands := []string{macOSCommandOption, command}
    return exec.Command(macOSCommand, commands...)
}

func (linux *Linux) buildCmd(command string) *exec.Cmd {
    commands := []string{linuxCommandOption, command}
    return exec.Command(linuxOSCommand, commands...)
}

func (windows *Windows) buildCmd(command string) *exec.Cmd {
    commands := []string{winCommandOption, command}
    return exec.Command(winOSCommand, commands...)
}

func doExecOSCmd(cmd *exec.Cmd) bool {
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
