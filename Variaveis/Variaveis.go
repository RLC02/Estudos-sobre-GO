package main

import (
	"fmt"
)

//Abrangecia de nivel de pacote.
//A variável aplha não pode ser acessada pelo mesmo nivel de pacote.
var aplha = "Girafa"

func main(){

	// := -> Marmota é utilizada para declaração e atribuição de variáveis.
	// = -> Atribuição de valor a variável já declarada.
	//Não se declara variavel com keyworlds reservadas.
	//func := 42 XX

	// Declaração de variáveis.
	x := 42
	y := "Marmota"

	// Atribuição de expressões a variáveis.
	v := 10 + 20

	// Atribuição de comparações a variáveis.
	w := 10 == 10
	i := 10 < 10

	// Imprimindo os valores das variáveis e seus tipos.
	fmt.Printf("x: %v, %T \n", x , x)
	fmt.Printf("y: %v, %T \n", y, y)

	fmt.Printf("v: %v \n\n", v,)

	fmt.Printf("w: %v \n", w,)
	fmt.Printf("i: %v \n\n", i,)

	// Atribuindo um novo valor a uma variável junto com a declaração de uma nova variável.
	x, z := 90, 40

	// Imprimindo os valores das variáveis e seus tipos.
	fmt.Printf("x: %v, %T \n", x , x)
	fmt.Printf("z: %v, %T \n", z , z)
	
}

func teste(){
		//fmt.Println(k) fora do codeblock essa variavel não é acessivel.

		//Essa variavel esta sendo declarada no nivel de pacote.
		fmt.Println(aplha)
}