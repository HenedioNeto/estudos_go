package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":"Henedio",
		"sobrenome":"Pedrosa",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome":{
			"primeiro":"Henedio",
			"segundo":"Bernardino",
		},
		"curso":{
			"primeiro":"Relações Internacionais",
			"segundo":"Analise e Desenvolvimento de Sistemas",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string {
		"nome":"Virgem",
	}
}