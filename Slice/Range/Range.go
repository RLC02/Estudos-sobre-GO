package main

import(
	"fmt"
)



func main(){

	//declarando com o tipo subjacente que é uma array
	slice := []string{"banana", "maça", "uva", "melancia", "abacaxi"}

	//utilizando o range para percorrer o slice
	for indice, valor := range slice{
		fmt.Println("No indice", indice, "temos o valor", valor)
	}
	
	//ele não permite a alteração do valor do slice sem o append
	//slice[5] = "larranja"
	slice = append(slice, "laranja")

	for _, valor := range slice{
		fmt.Printf("um dos valores desse slice é %v\n", valor)
	}
}