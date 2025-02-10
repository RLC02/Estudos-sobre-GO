package main

import(
	"fmt"
)



func main(){
	slice := make([]string, 26, 26)

	slice = append(slice, "Acre","Alagoas","Amapá","Amazonas","Bahia","Ceará","Espírito Santo","Goiás","Maranhão","Mato Grosso","Mato Grosso do Sul","Minas Gerais","Pará","Paraíba","Paraná","Pernambuco","Piauí","Rio de Janeiro","Rio Grande do Norte","Rio Grande do Sul","Rondônia","Roraima","Santa Catarina","São Paulo","Sergipe","Tocantins")

	fmt.Println(len(slice), cap(slice))

	for i := 0; i < len(slice); i++ {
		fmt.Printf("Index: %d, Value: %s\n", i, slice[i])
	}
	
}

//exercicio correto