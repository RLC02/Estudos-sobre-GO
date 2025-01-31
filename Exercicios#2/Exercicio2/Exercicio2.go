package main

import (
	"fmt"
)

func main() {

	x := 42 == 42
	y := 42 != 42
	z := 42 <= 42
	a := 42 < 42
	b := 42 >= 42
	c := 42 > 42

	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", x, y, z, a, b, c)
	
}

//exercicio correto 
