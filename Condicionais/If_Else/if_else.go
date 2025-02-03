package main

import (
	"fmt"
)

func main() {

	//se x for maior que 100, imprime X é maior que 100
	//se x for menor que 10, imprime X não é maior que 100
	//se x não e maior nem menor, imprime X não e maior nem menor, então é oq ?
	if x := 10; x >100 {
		fmt.Println("X é maior que 100")
	}else if x <10 { //é possivel utilizar quantas condições forem necessárias
		fmt.Println("X não é maior que 100")
	}else {
		fmt.Println("X não e maior nem menor, então é oq ?")
	}

}