package main

import (
	"runtime"
	"irocn.com/irocn-fcloud/cmd"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	cmd.Execute()
}
