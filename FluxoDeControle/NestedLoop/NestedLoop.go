
package main

import "fmt"

func main(){
	//loop aninhado
	//pode se utilizar quantos loops forem necessários, um dentro do outro.
	//Esse lopp vai de 0 a 12 horas e de 0 a 60 minutos.
	//a cada hora ele imprime 60 minutos.
	for horas := 1; horas <= 12; horas++ {

		//Imprime as horas
		fmt.Println("Horas: ", horas)
		//for para os minutos
		for minutos := 0; minutos <= 60; minutos++ {
			fmt.Printf("%d ", minutos)
		}
		fmt.Printf("\n")
	}
}