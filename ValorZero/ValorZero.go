package main

import (
	"fmt"
)


//O valor zero de um int é 0.
var x int
//O valor zero de um string é "".
var y string
//O valor zero de um bool é false.
var z bool
//O valor zero de um float é 0.0
var w float64

//O valor zero de pointers, functions, interfaces, channels. maps e slices é nil

func main(){

	fmt.Printf("x: %v, %T \n", x, x)
	fmt.Printf("y: %v, %T \n", y, y)
	fmt.Printf("z: %v, %T \n", z, z)
	fmt.Printf("w: %v, %T \n", w, w)

}