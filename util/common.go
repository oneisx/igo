package util

import (
    "strconv"
)

func IsNum(s string) bool {
    _, err := strconv.ParseInt(s, 0, 8)
    return err == nil
}

func ParseInt(s string) int {
    v, _ := strconv.ParseInt(s, 0, 8)
    return int(v)
}
