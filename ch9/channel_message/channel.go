package main

import (
	"fmt"
	"time"
)

func thrower(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
		fmt.Println("Thtrew >>", i)
	}
}

func cather(c chan int) {
	for i := 0; i < 5; i++ {
		num := <-c
		fmt.Println("Caught <<", num)
	}
}

func main() {
	c := make(chan int)
	go thrower(c)
	go cather(c)
	time.Sleep(100 * time.Millisecond)
}
