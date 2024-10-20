package main

import "fmt"

type usuario struct {
	nome string
	idade int8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Henedio"
	u.idade = 26
	fmt.Println(u)

	usuario2 := usuario{"Henedio", 26}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Henedio"}
	fmt.Println(usuario3)

}