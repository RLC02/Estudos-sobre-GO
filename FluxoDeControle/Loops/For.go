
package main

import "fmt"

func main(){

	//for é o único loop em Go, não tem while, do while, foreach.
	//for é usado para repetir um bloco de código várias vezes.
	//ele procura por uma condição, se for verdadeira, executa o bloco de código.
	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}

}