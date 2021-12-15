package util

import (
    "golang.design/x/clipboard"
)

func WriteText2Clipboard(str string) {
    clipboard.Write(clipboard.FmtText, []byte(str))
}
