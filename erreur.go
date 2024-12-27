package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(division(10452, 32424))
}

func division(a int , b float64) (float64) {
	if b == 0 {
		errors.New("on ne peut pas diviser par 0")
	} 
	return float64(a) / b
}
