package main

import (
	"fmt"
	"math"
)

// type comes after arguments; return type after prototype
func add(x int, y int) int {
	return x + y
}

// you can return pairs of ints, name return values, naked returns
func split(sum int) (x, y int) {
	x = sum / 2
	y = sum - x
	return
}

func main() {
	fmt.Printf("Now %g\n", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(split(10))
}
