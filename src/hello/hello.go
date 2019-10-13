package main

import (
	"fmt"
	"time"
	"math"
	"math/cmplx"
)

var c, python, java bool

func add(a int, b int) int {
	return a + b
}

func concat(a, b string) (string, string) {
	return a + b, b + a
}

func split(sum int) (x, y, z int) {
	z = 2
	x = 1
	y = sum - x
	return 
}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {	
	var i = 1
	aasd := 123
	fmt.Printf("hello, world\n")
	fmt.Printf("hello, world %g\n", time.Now())
	fmt.Printf("%g random number", math.Sqrt(10), "\n")
	fmt.Println(add(2,3))
	fmt.Println(concat("Hello","World"))
	fmt.Println(split(10))
	fmt.Println(i, c, python, java)
	fmt.Println(aasd)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)


	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}