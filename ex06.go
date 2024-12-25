package main

import ("fmt")

type Personnage struct {
	name string
	age int
	adresse string
}

func main() {
	var perso1 Personnage
	var perso2 Personnage

	perso1.name = "lucas"
	perso1.age = 21
	perso1.adresse = "1 rue de la pai"

}