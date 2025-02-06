package main

import(
	"fmt"
)



func main(){
	
	//aqui estamos criando um slice com o make 
	//o make permite que você crie um slice com um tamanho e capacidade definida do array subjacente
	//isso permite um gerenciamento mais eficiente da memória quando slice é utilizado
	slice := make([]int, 5, 10)

	slice[0], slice[1], slice[2], slice[3], slice[4] = 1, 2, 3, 4, 5

	slice = append(slice, 6, 7, 8, 9, 10)

	fmt.Println(slice, len(slice), cap(slice))

	//quando ele ultrapassa a capacidade do make ele deleta o array antigo e dobra a capacidade do array subjacente
	//isso pode ser um problema se você estiver trabalhando com um grande volume de dados
	//então é importante utilizar o make de maneira eficiente
	slice = append(slice, 10)

	fmt.Println(slice, len(slice), cap(slice))
	
}