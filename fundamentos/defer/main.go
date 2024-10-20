package main

import "fmt"

func funcao1() {
	fmt.Println("executando função1")
}

func funcao2() {
	fmt.Println("executando função2")
}

func main() {
	defer funcao1()
	//Adia a execução da funcao1 para ser a ultima executada
	funcao2()
	fmt.Println()
}