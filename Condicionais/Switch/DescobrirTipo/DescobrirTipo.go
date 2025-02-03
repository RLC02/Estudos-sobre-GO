package main

import (
	"fmt"
)

var x interface{}

func main() {

	x = true

	//é possivel descobrir o tipo de uma variavel utilizando o switch
	switch x.(type) {
	case int:	
		fmt.Println("x é um inteiro")
	case float64:
		fmt.Println("x é um float64")
	case string:
		fmt.Println("x é uma string")
	case bool:
		fmt.Println("x é um booleano")

	}
}