package main

import(
	"fmt"
)

func main(){
	
	hamburguer := map[string]int{	
		"X-Bacon": 15,
		"X-Salada": 10,
		"X-Tudo": 20,
		"X-Egg": 12,
	}

	fmt.Println(hamburguer)

	//aqui estamos imprindo todos o valores do map
	//o map possui um range aleatorio
	//então ele não imprime na ordem que foi criado
	for key, value := range hamburguer{
		fmt.Println(key, value)
	}

	//dessa maneira estamos deletando um valor do map
	delete(hamburguer, "X-Tudo")

	fmt.Println(hamburguer)

}