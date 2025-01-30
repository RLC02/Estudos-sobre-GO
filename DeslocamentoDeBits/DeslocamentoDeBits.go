package main

import (
	"fmt"
)

const (

	//exemplo pratico de deslocamento de bits
	KB = 1 << (10 * iota)
	MB 
	GB 
	TB 
)

func main() {

	x := 1
	//Deslocamento de bits
	y := x << 2
	z := x >> 2

	fmt.Printf("%b\n", x)
	fmt.Printf("%b\n", y)
	fmt.Printf("%b\n", z)
	
}