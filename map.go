package main
import ("fmt")

func main() {
	var tab = []int{1,2,3,4,5}
	var dict = map[string]int{"lucas":21,"emile":22,"theo":19,"nina":25}

	for i:=0; i < len(tab); i++ {
		fmt.Println(tab[i])
	}
	
	fmt.Println(dict)

}