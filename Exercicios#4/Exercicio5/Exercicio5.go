package main

import(
	"fmt"
)


func main(){

	x :=  []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	y := append(x, append(x[:3], x[len(x)-2:]...)...)

	fmt.Println(y)

}
//exercicio correto