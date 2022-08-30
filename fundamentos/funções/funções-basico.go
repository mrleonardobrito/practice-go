package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {
	fmt.Println(somar(3, 2))

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("função f")

	resultadosoma, resultadosubtracao := calculosMatematicos(3, 2)
	fmt.Println(resultadosoma, resultadosubtracao)
}
