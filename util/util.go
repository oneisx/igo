package util

import (
	"github.com/atotto/clipboard"
	"strconv"
)

func IsNum(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func Atoi(str string) int {
	v, _ := strconv.Atoi(str)
	return v
}

func WriteClipboard(str string) {
	_ = clipboard.WriteAll(str)
}
