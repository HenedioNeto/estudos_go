package main

import "fmt"

func main() {
	//Artiméticos
	soma := 10 + 5
	subtracao := 10 - 5
	divisao := 10 / 5
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 5

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	//Não da pra fazer operações com int de tipo diferente (int8 + int16)

	//Atribuição
	var variavel string = "String"
	variavel2 := "String2"
	fmt.Println(variavel, variavel2)

	//Operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	//Operadores lógicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//Operadores unarios
	numero := 10
	numero++
	fmt.Println(numero)
	numero += 1
	fmt.Println(numero)
}