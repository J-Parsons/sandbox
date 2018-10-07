package main

import "fmt"

func main() {
	var p *int
	v := 10
	p = &v
	fmt.Println(*p)
	// dereference with *, so *p is the value of v
	// unlike c, Go has no pointer arithmetic

	// a struct is a collection of fields
	// fields accessed with .
	type Point struct {
		X int
		Y int
	}
	cord := Point{1, 2}
	fmt.Println(cord.X)
	point := &cord
	point.Y = 10
	fmt.Println(point.Y)

	var (
		v1 = Point{1, 2}  // has type Point
		v2 = Point{X: 1}  // Y:0 is implicit
		v3 = Point{}      // X:0 and Y:0
		v4 = &Point{1, 2} // has type *Point
	)
	fmt.Println(v1, v2, v3, v4)

	//arrays
	var arr [5]int
	arr[0] = 0
	arr[1] = 1

	// unicode makes me happy and so does type inferencing
	// short assignment needs a type declaration when working with arrays
	次 := [4]string{"春", "夏", "秋", "冬"}
	fmt.Println(次)

	// slice. Or a structure containing pointers to elements of an array.
	// useful when you're not allowed to do pointer arithmetic
	無礼 := 次[2:4]
	fmt.Println(無礼)
	// I also learned about stacking defers
	b := 次[0:2]
	defer fmt.Println(b)
	b[0] = "面白い"
	defer fmt.Println(b)
	// trick question: line 50 messes with defer print from line 49
	// slices always stay up to-date because they refer to an address
	// even if it's deferred, printing a slice will reflect later updates

	// array literal: [3]bool{true, true, true,}
	// slice literal: []bool{true, true, true}
	// creates the same array as above then builds a reference to it
	// i.e. you don't need to have an array before a slices

	// length of slice = number of elements it contains,len(s)
	// capacity = number of elements in underlying array,cap(s)
	// you can modify a slice's length by slicing the slice

	s := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(s)

	s = s[:0]
	fmt.Println(s)

	s = s[:4]
	fmt.Println(s)

	s = s[2:]
	fmt.Println(s)

	// to create dynamic arrays, use make
	// a := make([]int, 5)
	// b := make([]int, 0, 5)

	// slices can contain ANY type (atomics, other slices...)

	// append handles all cases of adding elements to a slice
	s = append(s, 0)
	fmt.Println(s)

	s = append(s, 10, 20, 30, 40)
	fmt.Println(s)

	// python-style range for loops
	// scans the slice / map / etc by index
	for i := range s {
		fmt.Printf("%v: %v\n", i, s[i])
	}
	fmt.Println()

	// include value right away
	for i, value := range s {
		fmt.Printf("%v: %v\n", i, value)
	}
	fmt.Println()

	// you can drop the index or value with "_"
	for _, value := range s {
		fmt.Println(value)
	}
	fmt.Println()

	var m = map[string]int{
		"key1": 1,
		"key2": 2,
		"hi":   10,
	}
	fmt.Println(m)

	m2 := make(map[string]int)
	m2["key"] = 10

	// test that a key is present with a two-value assignment
	elem, ok := m2["key"]
	// if "key" is in m2, then ok will be true, else false
	// if "key" is not in the map, then elem is the zero value for map elem type
	fmt.Printf("elem: %v	ok: %v\n", elem, ok)
	elem, ok = m2["not here"]
	fmt.Printf("elem: %v	ok: %v\n", elem, ok)
}
