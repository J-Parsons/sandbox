package main

import (
	"fmt"
	"math/cmplx"
)

// factored variable declaration
var (
	check bool       = true
	z     complex128 = cmplx.Sqrt(-5 + 12i)
	// byte, rune, uint, sized int, float...
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", check, check)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var i int = 42
	// type conversion; type inference
	var j string = string(i)
	fmt.Println(i, j)

	// constant declared using "const" instead of "var"
	// style = camelcase
	// cannot be declared using := syntax
	const World = "世界"
	fmt.Println("hello", World)

}
