package main

import "fmt"

func printType(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Println("String value:", v)
	case int:
		fmt.Println("Integer value:", v)
	case bool:
		fmt.Println("Boolean value:", v)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	printType("hello") // ➜ String
	printType(42)      // ➜ Integer
	printType(true)    // ➜ Boolean
	printType(3.14)    // ➜ Unknown type
}
