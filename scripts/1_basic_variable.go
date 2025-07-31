package main

import "fmt"

func main() {

	var a int = 42
	var b int8 = -100
	var c int16 = 32000
	var d int32 = -2000000000
	var e int64 = 9223372036854775807

	var f uint = 100
	var g uint8 = 255
	var h byte = 65
	var i uint16 = 65535
	var j uint32 = 4294967295
	var k uint64 = 18446744073709551615

	var l float32 = 3.14159
	var m float64 = 2.718281828459045

	var n complex64 = complex(1.5, 2.5)
	var o complex128 = 2 + 3i

	var p rune = '₹'
	var q bool = true
	var r string = "Hello, Go!"

	fmt.Printf("a: %d, type: %T\n", a, a)
	fmt.Printf("b: %d, type: %T\n", b, b)
	fmt.Printf("c: %d, d: %d, e: %d\n", c, d, e)

	fmt.Printf("f: %d, g: %d, h as byte: %c, i: %d, j: %d, k: %d\n", f, g, h, i, j, k)

	fmt.Printf("l: %f, m: %f\n", l, m)
	fmt.Printf("n: %v, o: %v\n", n, o)

	fmt.Printf("p: %c, Unicode: %U\n", p, p)

	fmt.Printf("q: %t\n", q)
	fmt.Printf("r: %s, len: %d\n", r, len(r))
}
