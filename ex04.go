package main

import ("fmt")

func main() {
	fmt.Println(add(344, 56))

}

func add(a int, b int) int{
	
	if a > b {
		return a
	} else {
		return b
	}
}