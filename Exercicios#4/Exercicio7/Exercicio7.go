package main

import(
	"fmt"
)

func main(){

	ss := [][]string{{"James", "Bond", "Tiro ao alvo"},
	 {"Miss", "Moneypenny", "Gastar dinheiro"},
	 {"Dr.", "No", "Ser malvado"},
	}


	for _, v := range ss{
			fmt.Println(v[0])
		for _, v2 := range v{
			fmt.Printf("\t%v\n", v2)
		}
	}



}

//exercicio correto