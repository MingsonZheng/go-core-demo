package main

import (
	"048程序性能分析基础/common"
	"048程序性能分析基础/common/op"
	"errors"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
)

var (
	profileName    = "memprofile.out"
	memProfileRate = 8
)

func main() {
	f, err := common.CreateFile("", profileName)
	if err != nil {
		fmt.Printf("memory profile creation error: %v\n", err)
		return
	}
	defer f.Close()
	startMemProfile()
	if err = common.Execute(op.MemProfile, 10); err != nil {
		fmt.Printf("execute error: %v\n", err)
		return
	}
	if err := stopMemProfile(f); err != nil {
		fmt.Printf("memory profile stop error: %v\n", err)
		return
	}
}

func startMemProfile() {
	runtime.MemProfileRate = memProfileRate
}

func stopMemProfile(f *os.File) error {
	if f == nil {
		return errors.New("nil file")
	}
	return pprof.WriteHeapProfile(f)
}
