package main

import "fmt"

func main() {

	func(texto string) {
		fmt.Println(texto)
	}("Parametro da função")
	

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Parametro da função")

	fmt.Println(retorno)
}