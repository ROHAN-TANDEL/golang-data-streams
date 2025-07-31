package main

import "fmt"

var globalVar = "Visible everywhere in this package"

func main() {
	localVar := "Only in main"
	if true {
		blockVar := "Only inside if block"
		fmt.Println(globalVar, localVar, blockVar)
	}
	// fmt.Println(blockVar) // ❌ Error: undefined
}
