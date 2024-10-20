package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = 1000000000000000
	fmt.Println(numero)

	//a variavel iniciada sem valor vale 0
	var numero0 int16
	fmt.Println(numero0)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}

//alias:

//INT32 = RUNE
//UINT8 = BYTE

//int
//float
//bool
//error

