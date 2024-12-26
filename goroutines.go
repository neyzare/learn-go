package main

import (
	"fmt"
	"time"
)

func main() {
	go toto("goroutine")
	toto("direct")
}

func toto(from string) {
	for i:= 0; i<3; i++ {
		fmt.Println(from, ":", i)
	}
}