package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	for _, v := range a {
		a[3] = 10
		fmt.Println("Pass 1:", v)
	}
	// reset our mutation
	a[3] = 4
	// loop via pointer:
	b := &a
	for _, v := range b {
		b[3] = 10
		fmt.Println("Pass 2:", v)
	}
}
