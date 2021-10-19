package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse() // 把用户传递的命令行参数解析为对应变量的值
	fmt.Printf("Hello, %s!\n", name)
}

// go run demo2.go -name="Robert"
