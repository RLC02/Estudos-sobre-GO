package main

import (
	"fmt"
)

func main() {

	s := "ascii, éâ漢"
	//converte string para slice de bytes
	sb := []byte(s)

	//itera sobre a slice de bytes
	for _, v := range sb {
	//imprime o valor, o tipo, o unicode e o hexadecimal
	fmt.Printf("%v- %T - %U - %x\n", v, v, v, v)
	}

	for i:=0; i < len(s); i++ {
		fmt.Printf("%v - %T - %U - %x\n", s[i], s[i], s[i], s[i])
	}

}


