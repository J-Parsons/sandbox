package main

// go does not have classes, but you can define methods for types

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

// this is a method bound to a type
// in between func and the function name is receiver; type that receives method
// you can declare a method for any type defined in the same package
// this discludes built-in atomics like int

func (p Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// pointer receivers are common; allow value modification
// otherwise you would be working with a copy of the object
// remember that (*p).X is cumbersome, so Go lets us write p.X
func (p *Point) Scale(f float64) {
	p.X *= f
	p.Y *= f
}

// methods have pointer indirection for conveniance, but functions do not
// in general, all methods on a given type should have either value or pointer
// receievers, but not a mixture of both

// interfaces are types for polymorphism
// this way all the variants of a method are covered by a single type

// interfaces are (value, type) tuples
// executes the method of the same name on the underlying type
type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	p := Point{3, 4}
	fmt.Println(p.Abs())

	// interface type assertions
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}
