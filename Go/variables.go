package main

import "fmt"

var c, python, java bool

func main() {
	// var used to declare 1+ variables
	var i, j, k = 10, false, "hi"
	// if initialized, the value ca be ommitted!
	var or, well = "alternative", "solution"
	fmt.Println(i, c, python, java)
	fmt.Println(j, k, or, well)

	// short assignment: can be used in place of var with implicit typing
	// only works inside functions
	d := 10
	binary, hex := 2, 16
	fmt.Println(binary, hex, d)
}
