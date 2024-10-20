package main

import "fmt"

func main() {

	numero := 10

	if numero > 15{
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	//if init
	//A variavel é criada e serve apenas no escopo do if
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero maior que 0")
	}
}