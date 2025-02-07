package main

import(
	"fmt"
)

func main(){
	//é possivel ter quantas dimensões você quiser
	//aqui estamos criando um slice multidimensional 
	//eles são basicamente como planilhas do excel ou uma matriz
	ss := [][]int{
		//[]int{1,2,3,4,5,6},
		//[]int{6,7,8,9,10},
		//[]int{11,12,13,14,15},
	}

	//ele imprime o y que é igual a 2 e o x que é igual a 4
	//logo ele imprime o valor 17 que esta na linha 4 e coluna 2
	fmt.Println(ss[2][4])
}