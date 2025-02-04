
package main

import "fmt"

func main(){

		for i := 65; i <= 98; i++ {

		fmt.Println(i)
		for k := 0; k < 3; k++ {
			fmt.Printf("\t%#U\n", i)
		}

	}
}

//exercicio correto