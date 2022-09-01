package main

import "fmt"

func soma(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}

func imprimeTxtNumeros(texto string, numeros ...int) {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	fmt.Println(texto, total)
}

func main() {
	total := soma(1, 2, 3, 5, 5, 6, 7, 9, 64, 23, 34)
	imprimeTxtNumeros("Olá mundo", 212, 123, 12, 1, 23)
	fmt.Println(total)
}
