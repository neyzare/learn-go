package main

import ("fmt")

func main() {
	var count int = 120

	for i := 0; i <= count; i++ {
		
		if i % 3 == 0 {
			fmt.Println(i)
		}

	}
}