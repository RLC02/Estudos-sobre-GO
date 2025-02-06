package main

import(
	"fmt"
)

func main(){
	
	//aqui estamos criando um map
	//com keys e values
	marmotas := map[string]int{
		"Raul": 1,
		"Ricardo": 2,
		"Rafael": 3,
		"ruan": 4,
	}

	//aqui ele imprime o map inteiro e o valor da key Raul
	fmt.Println(marmotas)
	//aqui ele imprime o valor do Raul
	fmt.Println(marmotas["Raul"])

	//aqui estamos adicionando um novo valor ao map
	marmotas["gopher"] = 5

	//sera que o gasparzinho esta no map?
	//o ok é um booleano que retorna true se a key esta no map
	//e false se não esta
	//será, ok := marmotas["gasparzinho"]
	//fmt.Println(será, ok)

	//é possivel utilizar o if para verificar se a key esta no map
	//e imprimir o valor da key
	if será, ok := marmotas["gasparzinho"]; ok{
		fmt.Println(será)
	}else{
		fmt.Println("Não esta no map")
	}

}