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

	fmt.Println(slice[:3])
	fmt.Println(slice[4:])
	fmt.Println(slice[1:7])
	fmt.Println(slice[2:9])
	fmt.Println(slice[2:len(slice)-1])
	
}
//exercicio correto