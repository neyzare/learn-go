package main

import (
	"os"
	"fmt"
)

func main() {
	fichier, err := os.OpenFile("data.txt", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0600)
	defer fichier.Close()

	if err != nil {
		panic(err)
	}

	_,err = fichier.WriteString("hello world !\n")
	if err != nil {
		panic(err)
	}

	data,err := os.ReadFile("data.txt")
	fmt.Println(string(data))
}

