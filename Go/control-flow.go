package main

import (
	"fmt"
	"math"
	"runtime"
)

func foo() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
	fmt.Println(runtime.GOOS)
}

// if statements can begin wtih short statements like loops do
// declared values in control flow are removed once you leave scope
// vars declared in if short statements also avilable in subsequent else blocks
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%v >= %v\n", v, lim)
	}
	// but not here
	return lim
}

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// init and post statement are optional
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// no while loops
	con := true
	for con {
		sum += 10
		if sum > 1000 {
			break
		}
	}
	fmt.Println(sum)

	foo()
}
