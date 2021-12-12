package util

import (
    "os"
    "os/exec"
    "runtime"
)

const (
    winOSCommand       = "cmd.exe"
    winCommandOption   = "/c"
    winClearCommand    = "cls"
    linuxOSCommand     = "/bin/bash"
    linuxCommandOption = "-c"
    linuxClearCommand  = "clear"
    macOSCommand = "/usr/bin/open"
    macOSCommandOption = "-a"

)

func ClearScreen() {
    var cmd *exec.Cmd
    if isWindows() {
        cmd = buildWindowsCmd(winClearCommand)
    } else {
        cmd = buildLinuxCmd(linuxClearCommand)
    }
    doExecOSCmd(cmd)
}

func ExecOSCmd(command string) {
    cmd := buildCmd(command)
    doExecOSCmd(cmd)
}

func buildCmd(command string) *exec.Cmd {
    var cmd *exec.Cmd
    if isWindows() {
        cmd = buildWindowsCmd(command)
    } else if isLinux() {
        cmd = buildLinuxCmd(command)
    } else {
        cmd = buildMacOSCmd(command)
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

func isWindows() bool {
    return runtime.GOOS == "windows"
}

func isLinux() bool {
    return runtime.GOOS == "linux"
}
