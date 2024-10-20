package main

import "fmt"

func main() {
	var variavel1 string = "Variavel 1"
	//declaração implicita (inferencia de tipo)
	variavel2 := "Variavel2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	/*
	var (
		var variavel3 string = "variavel4"
		variavel4 := "Variavel5"
	)

	variavel5, variavel6 := "variavel5", "variavel6"

	const constante1 string = "constante1"

	inverte o valor das variaveis:
	variavel5, variavel6 = variavel6, variavel5
	*/
}