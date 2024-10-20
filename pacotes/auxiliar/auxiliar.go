package auxiliar

import "fmt"

//Para uma função ser lida fora do pacote ela deve começar com letra maiuscula.
//o escrever2 esta em letra minuscula, porém pode ser exportado para outros pacotes por estar dentro da função Escrever()
func Escrever() {
	fmt.Println("Escrevendo do modulo")
	escrever2()
}