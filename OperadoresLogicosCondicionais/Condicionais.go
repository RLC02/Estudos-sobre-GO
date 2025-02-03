package main

import (
	"fmt"
)


func main() {

	x := 2

	//o if pode ser utilizado para verificar se uma condição é verdadeira
	//utilizando o operador logico || ou
	 if (x == 2) || x == 3 {
		fmt.Println("x é 2 ou 3")
	 }

	 //ele retorna verdadeiro se as duas condições forem verdadeiras
	 //utilizano o operador logico && e
	 if (x %2 != 0 && x %3 == 0) {
		fmt.Println("x é 2 ou 3")
	 }
	
}