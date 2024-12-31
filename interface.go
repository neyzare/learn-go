package main 

import(
	"fmt"
)

func main() {
	var r Rectangle = Rectangle{5,5}	
	
	fmt.Println(r.area())
}

type Shape interface {
	area() float64
}

type Rectangle struct {
	largeur float64
	hauteur float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.hauteur * r.largeur 
}

func (c Circle) area() float64 {
	return c.radius	
}