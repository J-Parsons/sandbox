/*  https://tour.golang.org/methods/20
Update Sqrt from https://tour.golang.org/flowcontrol/8 to return an error value
*/
package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	delta, z := 0., 1.
	for {
		delta, z = z, z-(z*z-x)/(2*z)
		if math.Abs(delta-z) < 1e-8 {
			break
		}
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
