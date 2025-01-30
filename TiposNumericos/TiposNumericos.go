package main

import (
	"fmt"
)

//declarando variavel do tipo booleano
var x bool

func main() {
	//1 BYTE 
	a := "e"
	//2 BYTE 
	b := "é"
	//3 BYTE 
	c := "漢"
	fmt.Printf("%v, %v, %v\n", a, b, c)

	d := []byte("a")
	e := []byte("b")
	f := []byte("c")

	fmt.Printf("%v, %v, %v\n", d, e, f)

}



