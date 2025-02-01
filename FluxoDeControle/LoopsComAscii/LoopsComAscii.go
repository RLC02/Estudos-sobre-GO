
package main

import "fmt"

func main(){

	//Loop de 33 a 122, que são os valores da tabela ASCII.
	for i := 33; i <= 122; i++ {

		//Imprime o valor de i, o valor de i em Unicode e o valor de i em Unicode.
		fmt.Printf("%v\n", string(rune(i)))
		
	}

}

//exercicio correto