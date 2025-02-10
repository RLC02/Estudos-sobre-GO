package main

import(
	"fmt"
)

func main(){

	mapinha := map[string][]string{
		"Junior_James": {"jogar bola", "estudar", "trabalhar"},
		"Silva_Jaqueline": {"cozinhar", "filmes", "dormir"},
		"Marx_Karl ": {"comunismo", "socialismo", "escrever"},
	}
	
	mapinha["Blade_Sonia"] = []string{"matar vampiros", "chutar bundas", "fazer justiça"}

	for k, v := range mapinha{
		fmt.Println(k)
		for i, v2 := range v{
			fmt.Printf("\t%v - %v\n", i, v2)
		}
	}

}

//exercicio correto