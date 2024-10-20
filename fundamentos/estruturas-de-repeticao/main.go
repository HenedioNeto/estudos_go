package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando 1")
		time.Sleep(time.Second)
	}

	fmt.Println(i)

	for j := 0; j < 10; j++{
		fmt.Println(j)
	}

	numeros := [3]int{1, 2, 3}

	for indice, numero := range numeros {
		fmt.Println(indice, numero)
	}

	for _, numero := range numeros {
		fmt.Println(numero)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}
}