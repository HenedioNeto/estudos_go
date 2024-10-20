package main

import "fmt"

//Não é bem uma herança, apenas o mais proximo disso

type pessoa struct {
	nome string
	sobrenome string
	idade uint8
	altura uint8
}

type estudante struct {
	pessoa
	curso string
	faculdade string
}

func main() {
	pessoa1 := pessoa{"Henedio", "Neto", 26, 177}
	fmt.Println(pessoa1)

	estudante1 := estudante {pessoa1, "ADS", "FATEC"}
	fmt.Println(estudante1)
}