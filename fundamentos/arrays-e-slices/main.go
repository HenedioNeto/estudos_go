package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array1[5] int
	array1[0] = 1
	fmt.Println(array1)

	array2 := [3]string{"Posição1", "Posição2", "Posição3"}
	fmt.Println(array2)
	//Tamanho fixo e dados fixos

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{10, 11, 12, 12, 13, 14}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18)
	fmt.Println(slice)

	//Cria um slice a partir de um array existente
	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição alterada"
	fmt.Println(slice2)

	//Arrays internos
	//Função make aloca um espaço na memória
	fmt.Println("-----------")
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //Tamanho
	fmt.Println(cap(slice3)) //Capacidade

	//Quando o Go ve que a capacidade do slice vai ser ultrapassada ele aumenta a capacidade do slice
}