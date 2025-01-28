package main

import (
	"fmt"
)

//tipos primitivos int, string, bool
//tipos compostos array, slice, map, struct
//A declaração pega o tipo automaticamente.
var aplha = "Girafa"

func main(){

	//tipos são estaticos e não dinamicos.
	//Não podem ser alterados futuramente.
	//O compilador não aceita a atribuição de um valor de um tipo diferente do tipo da variável.
	//Alpha = 34 xx
	fmt.Printf("alpha: %v, %T \n", aplha, aplha)
}