package main

import (
	"003库源码文件/demo5/lib"
	"flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	lib.Hello(name)
}

// go env -w GO111MODULE=off
