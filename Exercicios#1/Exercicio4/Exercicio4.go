package main

import (
	"fmt"
)

type Prova int
var x Prova 

func main() {

	fmt.Printf("x: %v, %T \n", x, x)
	x = 42
	fmt.Println("x:", x)
}

	//exercicio correto 