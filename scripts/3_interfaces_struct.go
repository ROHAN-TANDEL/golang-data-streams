package main

import (
	"fmt"
	"reflect"
)

// ------------------------------
// INTERFACE DEFINITION
// ------------------------------
type Speaker interface {
	Speak() string
}

// Struct with function field and interface field
type Dog struct {
	Name    string
	OnSpeak func() string // function field
	Voice   Speaker       // interface field
}

// Struct implementing the interface
type Cat struct {
	Name string
}

// Dog implements the Speaker interface
func (d Dog) Speak() string {
	return "Woof! I am " + d.Name
}

// this is like registering a function to the interface
func (c Cat) Speak() string {
	return "Meow! I am " + c.Name
}

func main() {
	s, c := Speaker(Dog{Name: "Max"}), Speaker(Cat{Name: "Whiskers"})
	fmt.Println("Interface Example:")
	fmt.Println(s.Speak())
	fmt.Println("Type of s:", reflect.TypeOf(s))

	fmt.Println("Interface Example:")
	fmt.Println(c.Speak())
	fmt.Println("Type of s:", reflect.TypeOf(c))
}
