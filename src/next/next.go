package main

import (
	"fmt"
)

type User struct {
	name     string
	username string
}

var mapUser = map[string]User{
	"hello": User{
		name:     "nametext",
		username: "usernametext",
	},
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func main() {
	// mapUser = make(map[string]User)
	// mapUser["hello world"] = User{"vladimir", "saakyan"}
	// mapUser["hello world2"] = User{"vladimir", "saakyan"}
	for i, v := range mapUser {
		fmt.Println(i, v)
	}
	fmt.Println(mapUser)
	var sc int = 1
	fmt.Println(sc)
}
