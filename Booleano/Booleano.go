package main

import (
	"fmt"
)

//declarando variavel do tipo booleano
var x bool

func main() {
	fmt.Println(x)//valor zero
	x = true
	fmt.Println(x)//valor atribuido
	//operadores relacionais
	x = 10 < 100 //menor
	y := 10 == 100//igual
	z := 10 > 100//maior
	fmt.Println(x, y, z)
}


