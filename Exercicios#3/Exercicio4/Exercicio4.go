
package main

import "fmt"

func main(){

	anon := 2007
	anoq := 2025

	for{
		if anon > anoq{
			break
		}
		fmt.Println(anon)
		anon++
	}

}
//exercicio correto