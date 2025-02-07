package main

import(
	"fmt"
)


func main(){

	x := [5]int{1, 10, 20, 30, 40}

	for i, v := range x{
		fmt.Printf("No indice: %v temos o valor: %v\n", i, v)
	}
	fmt.Printf("O tipo de x é %T\n", x)
	
}
//exercicio correto