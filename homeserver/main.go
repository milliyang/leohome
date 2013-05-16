package main

import (
	"fmt"
	"github.com/axgle/pinyin"
)

func main() {
	fmt.Println(pinyin.Convert("hello,百度新浪网易!"))
}
