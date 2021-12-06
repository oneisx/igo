package biz

import (
	"fmt"
	"time"
)

//func Transfer(args []string) []string {
//	if len(args) == 0 {
//		return nil
//	}
//	return nil
//}

func main() {
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixMilli())
	fmt.Println(time.Now().UnixMicro())
	fmt.Println(time.Now().UnixNano())
}
