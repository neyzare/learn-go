package main

import (
	"errors"
	"fmt"
)

func main() {
	
	resultat, err := division(10452, 43523)

	if err != nil {
		fmt.Println("erreur", err)
	} else {
		fmt.Println("le resultat est : ", resultat)
	}
}

func division(a int , b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("on ne peut pas diviser par 0")
	} 
	return float64(a) / b, nil
}
