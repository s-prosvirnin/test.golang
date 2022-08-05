package main

import (
	"fmt"
	"time"
)

func main() {
	f1()
	time.Sleep(time.Second * 10)
}

func f1() {
	go func() {
		for {
			fmt.Println("go1")
			time.Sleep(time.Millisecond * 500)
		}
	}()
}
