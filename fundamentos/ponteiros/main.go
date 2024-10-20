package main

import "fmt"

//Aponta um enderço de memória para a variavel

func main() {
	var variavel int = 10
	var variavel2 int = variavel

	fmt.Println(variavel, variavel2)

	variavel++
	fmt.Println(variavel, variavel2)
	//Apenas o valor da primeira variavel muda

	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, *ponteiro) //Desreferenciação
}