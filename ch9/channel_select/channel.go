package main

import (
	"fmt"
	"time"
)

func callerA(c chan string) {
	c <- "Hello World!"
	close(c)
}

func callerB(c chan string) {
	c <- "Hola Mundo!"
	close(c)
}

func main() {
	a, b := make(chan string), make(chan string)
	go callerA(a)
	go callerB(b)

	var msg string
	ok1, ok2 := true, true
	for ok1 || ok2 {
		time.Sleep(1 * time.Microsecond)
		select {
		case msg, ok1 = <-a:
			if ok1 {
				fmt.Printf("%s from A\n", msg)
			}

		case msg, ok2 = <-b:
			if ok2 {
				fmt.Printf("%s from B\n", msg)
			}
		}
	}
}
