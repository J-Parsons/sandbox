package main

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, val := range s {
		sum += val
	}
	c <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int)
	go sum(s[len(s)/2:], c)
	go sum(s[:len(s)/2], c)
	x, y := <-c, <-c // received from channel c
	fmt.Println(x + y)
}
