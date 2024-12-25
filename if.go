package main
import ("fmt")

func ex02() {
	var nombre int
	fmt.Print("entre un nombre pair ou impair : ")
	fmt.Scan(&nombre)

	if nombre % 2 == 0 {
		fmt.Println("votre nombre est pair")
	} else if nombre % 2 != 0 {
		fmt.Println("votre nombre est impair")
	}

}

func main()  {
	ex02()
}