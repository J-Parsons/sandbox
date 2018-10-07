/*  https://tour.golang.org/concurrency/8
check binary tree equivalence using the tree package
type Tree struct {
	Left *Tree
    Value int
    Right *Tree
}
1. Implement the Walk function
2. Implement the Same function using Walk to determine whether t1 and t2 tore the same values
*/
package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree sending all valuas from the tree to channel ch
func Walk(t *tree.Tree, c chan int) {
	// recursion handler
	walkRecHandler(t, c)
	// close the channel so range will stop
	close(c)
}

// recurse down the tree
// recurse left, send the current node value to the channel, then recurse right
func walkRecHandler(t *tree.Tree, c chan int) {
	if t != nil {
		walkRecHandler(t.Left, c)
		c <- t.Value
		walkRecHandler(t.Right, c)
	}
}

// determines whether the trees pointed to by t1 and t2 are the same
func Same(t1, t2 *tree.Tree) bool {
	c1, c2 := make(chan int), make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)

	for {
		// xx gets a value from the channel
		// okx is T/F depending on whether a value was present or not
		x1, ok1 := <-c1
		x2, ok2 := <-c2
		switch {
		case ok1 != ok2:
			// trees run out of values at different times -> not same size
			return false
		case !ok1:
			// if the switch statement makes it here, we know the trees
			// are currently the same size
			// if !ok1, then both trees are now empty w/o encountering mismatch
			return true
		case x1 != x2:
			// an mismatch between elements means the trees aren't equivalent
			return false
		default:
			// do nothing; keep iterating
		}
	}

}

func main() {
	c := make(chan int)
	go Walk(tree.New(1), c)
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
