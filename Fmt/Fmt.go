package main

import (
	"fmt"
)

func main(){

	//string interpretada(Interpreted) exibe como um texto normal.
	x := "Olá, marmota\n Como vai você?\t eu \"vou\" bem."

	//string crua(raw) exibe como está escrito.
	//x = `"Olá, marmota\n Como vai você?\t eu \"vou\" bem."`

	//print com quebra de linha.
	fmt.Println(x)
	//print sem quebra de linha.
	//Sprintf não joga o valor na tela, armazena.
	//Fprintf joga o valor em um arquivo.

}