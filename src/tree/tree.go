package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		fmt.Println(v)
		sum += v
	}
	c <- sum
}

func main() {
	// s := []int{7, 2, 8, -9, 4, 0}

	// c := make(chan int, 2)

	// go sum(s[:len(s)/2], c)
	// go sum(s[len(s)/2:], c)
	// x, y := <-c, <-c // receive from c

	// fmt.Println(x, y)
	// fmt.Printf("%v", c)

	ch := make(chan int, 100)
	ch <- 20
	ch <- 10
	ch <- 30
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	close(ch)
}
