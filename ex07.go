package main

import (
	"fmt"
)

func main() {
	var pointeurString *string
	var lucas string = "lucas"
	pointeurString = &lucas

	fmt.Println("adresse de votre variable lucas  : ", &lucas)
	fmt.Println("valeur de la variable (pointeur) pointeurString %p\n ", pointeurString)
	fmt.Println("valeur de l'adresse %p: %d\n", pointeurString, *pointeurString)

}