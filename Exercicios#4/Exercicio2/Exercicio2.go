package main

import(
	"fmt"
)


func main(){

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range slice{
	fmt.Printf("No indice: %v temos o valor: %v\n", i, v)
	}

	fmt.Printf("O tipo de slice é %T\n", slice)
	
}
//exercicio correto