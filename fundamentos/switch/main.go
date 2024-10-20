package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3: 
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7: 
		return "Sabado"	
	default:
		return "Numero invalido"
	}
}

func main() {
	dia := diaDaSemana(3)
	fmt.Println(dia)
}