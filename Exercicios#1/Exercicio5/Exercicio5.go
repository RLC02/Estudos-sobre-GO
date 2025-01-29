package main

import (
	"fmt"
)

type Prova int
var x Prova 

var y int

func main() {

	fmt.Printf("x: %v, %T \n", x, x)
	x = 42
	fmt.Println("x:", x)

	y = int(x)

	fmt.Printf("y: %v, %T", y, y)
}

	//exercicio correto 