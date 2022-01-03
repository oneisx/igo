package cmd

import (
    "fmt"
    "igo/util"
)

type BaseCmd interface {
    exec(cmd string, args []string)
}

type ChangeDirectory struct {}

type PrintWorkDirectory struct {}

type ListDirectory struct {}

func (c *ChangeDirectory) exec(cmd string, args []string)  {
    util.Operator.ChangeDirectory(args[0])
}

func (p *PrintWorkDirectory) exec(cmd string, args []string)  {
    fmt.Println(util.Operator.PrintWorkDirectory())
}

func (l *ListDirectory) exec(cmd string, args []string)  {
    util.Operator.ListDirectory()
}