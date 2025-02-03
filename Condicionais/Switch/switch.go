package main

import (
	"fmt"
)

func main() {

	x := 10

	//switch é uma forma de simplificar o if else
	//não é necessário colocar o break, pois ele já é implicito
	//ele executa o primeiro case que for verdadeiro
	//é possivel utilizar quantos cases forem necessários
	//não existe fallthrough implicito
		switch{
		case x > 5:
			fmt.Println("x é maior que 5")
			//ela cai para o proximo case
			fallthrough
		case x < 5:
			fmt.Println("x é menor que 5")
		case x == 5, x == 10: //é possivel utilizar a codição ou
			fmt.Println("x é igual a 5")
		default: //caso nenhum case seja verdadeiro ele executa o default
		fmt.Println("Não tem nada, palhaçada")
	}
}