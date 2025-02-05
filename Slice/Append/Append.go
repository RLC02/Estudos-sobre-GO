package main

import(
	"fmt"
)



func main(){

	umaslice := []int{1,2,3,4}
	outraslice := []int{6,7,8,9,10}

	fmt.Println(umaslice)

	//aqui estamos adicionando valores adicionais ao slice com o append
	umaslice = append(umaslice, 5, 6, 7, 8)

	fmt.Println(umaslice)

	//aqui estamos adicionando um slice ao outro, é necessario utilizar o ... para desmembrar o slice
	//na documentação do go é chamado de enumeration 
	umaslice = append(umaslice, outraslice...)

	fmt.Println(umaslice)

}