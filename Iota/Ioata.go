package main

import (
	"fmt"
)

const	(

	//iota é uma constante que é incrementada a cada linha
	//a a = iota + 100 é possivel definir um valor inicial
	a = iota
	_ = iota
	c = iota
	//descarta o valor de iota
	_ = iota
	y = iota
	z = iota

)


func main() {

	 fmt.Println(a, c, y, z)
	
}