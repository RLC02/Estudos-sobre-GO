
package main

import "fmt"

func main(){

	for i := 0; i < 20; i++ {

		//Se o número for par, ele pula para o próximo número.
		//utilizando modulo para verificar se o número é par.
		//basicamente, se o resto da divisão por 2 for 0, o número é par.
		if i % 2 == 0 {
			//faz isso quando o número não é par
			//ele para a execução do loop atual e pula para o próximo número.
			//o break quebra a interação do loop, o continue quebra e pula para a próxima interação.
			continue
		}
		fmt.Println(i)
	}
}