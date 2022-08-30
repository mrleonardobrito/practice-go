package main

import "fmt"

func main() {
	// Estruturas de controles

	// IF

	numero := 10

	if numero < 10 {
		fmt.Println("Numero menor que 10")
	} else {
		fmt.Println("Numero maior que 10")
	}

	// IF com inicialização de variável

	// variável outronumero tem o escopo desse if
	if outronumero := numero; outronumero > 0 {
		fmt.Println("Outronumero maior que 0")
	} else {
		fmt.Println("Outronumero menor que 0")
	}
}
