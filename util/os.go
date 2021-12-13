package util

import (
    "fmt"
    "os"
    "os/exec"
    "runtime"
)

const (
    windows            = "windows"
    linux              = "linux"
    macos              = "darwin"
    winOSCommand       = "cmd.exe"
    winCommandOption   = "/c"
    winClearCommand    = "cls"
    linuxOSCommand     = "/bin/bash"
    linuxCommandOption = "-c"
    linuxClearCommand  = "clear"
    macOSCommand       = "/usr/bin/open"
    macOSCommandOption = "-a"
)

func ClearScreen() {
    var cmd *exec.Cmd
    switch runtime.GOOS {
    case windows:
        cmd = buildWindowsCmd(winClearCommand)
    case macos:
        fallthrough
    case linux:
        cmd = buildLinuxCmd(linuxClearCommand)
    default:
        fmt.Println("Error: Operation system is not supported!")
        os.Exit(1)
    }
    doExecOSCmd(cmd)
}

func ExecOSCmd(command string) {
    cmd := buildIgoCmd(command)
    doExecOSCmd(cmd)
}

func buildIgoCmd(command string) *exec.Cmd {
    var cmd *exec.Cmd
    switch runtime.GOOS {
    case windows:
        cmd = buildWindowsCmd(command)
    case linux:
        cmd = buildLinuxCmd(command)
    case macos:
        cmd = buildMacOSCmd(command)
    default:
        fmt.Println("Error: Operation system is not supported!")
        os.Exit(1)
    }
    return cmd
}

func buildMacOSCmd(command string) *exec.Cmd {
    commands := []string{macOSCommandOption, command}
    return exec.Command(macOSCommand, commands...)
}

func buildLinuxCmd(command string) *exec.Cmd {
    commands := []string{linuxCommandOption, command}
    return exec.Command(linuxOSCommand, commands...)
}

func buildWindowsCmd(command string) *exec.Cmd {
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
