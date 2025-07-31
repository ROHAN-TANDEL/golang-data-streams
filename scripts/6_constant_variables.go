package main

import "fmt"

func main() {

	const (
		A = iota // 0
		B        // 1
		C        // 2
	)

	const pi = 3.14                  // constant variable
	const greeting = "Hello, World!" // string constant

	fmt.Println("Value of pi:", pi)
	fmt.Println("Greeting message:", greeting)

	fmt.Println("Value of A:", A)
	fmt.Println("Greeting C:", C)
}
