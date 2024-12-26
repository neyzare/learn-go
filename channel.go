package main

import (
	"fmt"
)

func run(c chan string, name string) {
	c <- name
}

func main() {
	canal := make(chan string)
	go run(canal, "lucas")
	fmt.Println(canal)
}