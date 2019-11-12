package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"runtime"
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
	MaxInt uint64     = 0
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func splits() (a, b int) {
	a = 123
	b = 2
	return
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	v := math.Pow(x, n)
	if v < lim {
		fmt.Println(n)
		fmt.Printf("%g < %g\n", v, lim)
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	// world.printSomething()

	// fmt.Println(c)
	// fmt.Println(splits())
	// fmt.Println(sqrt(2), sqrt(47))
	// fmt.Println(
	// 	pow(3, 2, 10),
	// 	pow(3, 3, 20),
	// )
	fmt.Print("Go runs on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
	// var i = 1
	// aasd := 123
	// fmt.Printf("hello, world\n")
	// fmt.Printf("hello, world %v\n", time.Now())
	// fmt.Printf("%v random number", math.Sqrt(10))
	// fmt.Println(add(2, 3))
	// fmt.Println(concat("Hello", "World"))
	// fmt.Println(split(10))
	// fmt.Println(i, c, python, java)
	// fmt.Println(aasd)
	// fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	// var xx, yyy = fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)

	// fmt.Println(xx, yyy)
	// // var f float64
	// // var b bool
	// // var s string
	// // fmt.Printf("%v %v %q\n", f, b, s)

	// var x, y int = 3, 25
	// var f float64 = math.Sqrt(float64(x*x + y*y))
	// var z uint = uint(f)
	// fmt.Println(x, y, z)

	// v := 42 // change me!
	// w, t := fmt.Printf("v is of type %T\n", v)
	// fmt.Println(w, t)
}
