package main

import(
	"fmt"
)



func main(){

	//sequencia finita de elementos de um tipo
	array := [5]int{1,2,3,4,5}
	fmt.Println(array)
	//já o slice é um pedaço de um array
	//ele é dinamico e pode ser alterado
	slice := []int{1,2,3,4,5}
	fmt.Println(slice)

	//é possivel adicionar elementos ao slice
	//para isso é utilizado o append
	slice2 := append(slice, 6)
	fmt.Println(slice2)

	//é possivel alterar o valor de um elemento do slice
	fmt.Println(slice[3])
	slice[3] = 10

}