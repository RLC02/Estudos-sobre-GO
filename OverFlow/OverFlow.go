package main

import (
	"fmt"
)

func main() {

	var i uint16
	i = 65535
	//i = 65536 //erro limete do tipo uint16
	fmt.Println(i)
}