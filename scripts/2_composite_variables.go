package main

import "fmt"

func main() {
	// --------------------------------
	// 1️⃣ VARIABLE DECLARATION
	// --------------------------------

	var a int = 10        // explicit type and value
	var b = 20            // type inferred
	c := 30               // short declaration (within functions)
	var d string          // default zero value: empty string
	var e, f = 3.14, true // multiple variables

	fmt.Println("Variables:")
	fmt.Println("a:", a, "b:", b, "c:", c, "d:", d, "e:", e, "f:", f)
	fmt.Println()

	// --------------------------------
	// 2️⃣ COMPOSITE TYPES
	// --------------------------------

	// Array
	var arr = [3]int{1, 2, 3}

	// Slice
	slice := []string{"go", "is", "fun"}

	// Map
	m := map[string]int{"apple": 5, "banana": 3}

	// Struct
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Alice", Age: 25}

	// Pointer
	ptr := &a

	fmt.Println("Composite Types:")
	fmt.Println("Array:", arr)
	fmt.Println("Slice:", slice)
	fmt.Println("Map:", m)
	fmt.Println("Struct:", p)
	fmt.Println("Pointer to a:", *ptr)
	fmt.Println()

}
