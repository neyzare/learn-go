package main

import (
	"fmt"
	"time"
)

func main() {
	go toto("goroutine")
	toto("direct")

	go func() {
		fmt.Println("msg 1")
	}()
	go func() {
		fmt.Println("msg 2")
	}()

	time.Sleep(1 * time.Second)
}

func toto(from string) {
	for i:= 0; i<3; i++ {
		fmt.Println(from, ":", i)
	}
}
