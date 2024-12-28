package main

import ("fmt")

type Personnage struct {
	name string
	age int
	adresse string
}
func (p Personnage) great() {
	p.name = "alice"
	p.age = 25
	fmt.Println("Bonjour je m'appelle ", p.name, "et j'ai ", p.age,  "ans")
}

type Rectangle struct {
	width float64
	height float64
}

func (r Rectangle) air() float64 {
	 return r.width * r.height
}

func (r Rectangle) perimetre() (float64) {
	return 2 * r.width * r.height
}

func main() {
	var perso1 Personnage
	var perso2 Personnage

	perso1.name = "lucas"
	perso1.age = 21
	perso1.adresse = "1 rue de la paix"

	perso2.name = "nina"
	perso2.age = 23
	perso2.adresse = "1 rue napoleonsd"

	rectangle1 := Rectangle{width: 34,height: 53}
	rectangle2 := Rectangle{width: 543, height: 43}

	fmt.Println(perso1)
	fmt.Println(perso2)
	perso1.great()

	fmt.Println(rectangle1.air())
	fmt.Println(rectangle2.perimetre())

}

