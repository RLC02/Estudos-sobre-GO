package main

import (
	"fmt"
)

//criando um tipo sushi com base em int
type sushi int 
//criando uma variável do tipo sushi
var lanche sushi = 50

func main(){

	y := 20

	//os tipos são diferentes mesmo que usem a base do int.
	//variavel do tipo int
	fmt.Printf("y: %v ", y)
	//convertendo a variavel y para o tipo sushi
	lanche = sushi(y)
	//variavel do tipo sushi
	fmt.Printf("lanche: %v ", lanche)

}