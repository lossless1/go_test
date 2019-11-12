package main

import (
	"fmt"
	"strings"
)

type User struct {
	name      string
	surname   string
	years_old int
}

type BoolInt struct {
	i int
	b bool
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func main() {
	const default_symbol string = "_"
	board := [][]string{
		[]string{default_symbol, default_symbol, default_symbol},
		[]string{default_symbol, default_symbol, default_symbol},
		[]string{default_symbol, default_symbol, default_symbol},
	}

	for i := 0; i < len(board); i++ {
		fmt.Printf("%v\n", strings.Join(board[i], " "))
	}

	var arr []int
	arr = append(arr, 0)
	arr = append(arr, 0, 3, 4, 5, 6)
	// delete(arr, 2)
	s := "hello world"
	println(s)
	// arrf := append(arrs, 0, 2, 3)

	fmt.Println(arr)
	// The players take turns.
	// board[0][0] = "X"
	// board[2][2] = "O"
	// board[1][2] = "X"
	// board[1][0] = "O"
	// board[0][2] = "X"

	// for i := 0; i < len(board); i++ {
	// 	fmt.Printf("%s\n", strings.Join(board[i], " "))
	// }
}

// func main() {

// 	var user User = User{}
// 	var user2 User = User{"hello", "world", 123}

// 	p := &user
// 	p.name = "wow"
// 	fmt.Println(user)
// 	fmt.Println(user2)
// 	fmt.Println(*p)

// 	a := [5]int{0, 1, 2, 3, 4}
// 	var g []int = a[3:4]
// 	fmt.Println(a)
// 	fmt.Println(g)

// 	names := [4]string{
// 		"John",
// 		"Paul",
// 		"George",
// 		"Ringo",
// 	}
// 	fmt.Println(names)

// 	aa := names[0:2]
// 	b := names[1:3]
// 	fmt.Println(a, b)

// 	b[0] = "XXX"
// 	fmt.Println(aa, b)
// 	fmt.Println(names)
// }
