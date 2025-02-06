package main

import(
	"fmt"
)

func main(){
//
	sabores := []string{"queijo", "calabresa", "portuguesa", "marguerita", "atum"}

	//aqui estamos removendo o sabor portuguesa
	//deleteando o valor do slice
	//e definindo um range que vai de 0 até 2 e de 3 até o final
	sabores = append(sabores[:2], sabores[3:]...)

	//aqui estamos fatiando o slice
	//como fatiar uma fatia de pizza
	//ele recorta o valor apontado o incluindo
	fatia := sabores[2:5]
	fmt.Println(fatia)

	//Varias maneira de passar por todos os elementos de um slice sem utilizar o range
	fmt.Println(sabores[0], sabores[1], sabores[2], sabores[3], sabores[4])
	fmt.Println(fatia[0:4])

	for i := 0; i < len(sabores); i++{
		fmt.Println(sabores[i])
	}
	
}