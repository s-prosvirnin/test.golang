package main

import "fmt"

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		chan1 <- "11"
		chan1 <- "12"
		chan1 <- "13"
		fmt.Println("chan1 <-")
		close(chan1)
		fmt.Println("chan1 closed")
	}()

	go func() {
		chan2 <- "21"
		chan2 <- "22"
		chan2 <- "23"
		fmt.Println("chan2 <-")
		close(chan2)
		fmt.Println("chan2 closed")
	}()

	res1 := false
	res2 := false
	for !res1 || !res2 {
		select {
		case val, ok := <-chan1:
			fmt.Println("<-chan1: " + val)
			if !ok {
				res1 = true
			}
			//close(chan1)
		case val, ok := <-chan2:
			fmt.Println("<-chan2: " + val)
			if !ok {
				res2 = true
			}
			//close(chan2)
		default:
			break
		}
	}
}
