package main

import(
	"fmt"
)

//utilizado normalmente para alocar memorioa de maneira eficiente
//isso é um array de 5 posições
//um array nunca pode ser negativo
//ele é imutavel e tem um tamanho fixo
var x [5]int

func main(){

	//aqui estamos atribuindo valores para as posições do array
	x[0] = 1
	x[1] = 10

	//ele ira imprimir o valor das posições 0 e 1
	fmt.Println(x[0], x[1])
	//ele ira imprimir o valor da posição 0 como 1 e o valor da posição 1 como 10 e a restantes como 0 por padrão
	fmt.Println(x)
	//ele ira imprimir o tamanho do array
	fmt.Println(len(x))
}