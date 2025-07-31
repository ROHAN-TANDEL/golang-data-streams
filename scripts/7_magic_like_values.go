package main

import (
	"fmt"
	"runtime"
)

func main() {
	_, file, line, _ := runtime.Caller(0)
	fmt.Println("File:", file)
	fmt.Println("Line:", line)
	//output
	// File: /path/to/your/file/7_magic_like_values.go
	// Line: 9
}
