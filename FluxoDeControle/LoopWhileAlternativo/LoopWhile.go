
package main

import "fmt"

func main(){

	x := 0

	//while alternativo
	//não existe while em Go, mas é possível simular um while com o for
	//for x < 10 {
		//fmt.Println("X ainda é menor que 10")
		//x++
	//}

	//for com uma condicional if
	//um loop praticamente infinito, pois a condição é verdadeira
	for{
		if x < 10 {
			x++
			fmt.Println("X ainda é menor que 10")
		}
		fmt.Println("X ainda é maior que 10")
		//Porem o break consegue parar o loop e quebrar a execução
		break
	}
	fmt.Println("PAROUUUU")
}
