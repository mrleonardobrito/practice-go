package main

import "fmt"

// retorno nomeado
func calculosMatematicos(n1, n2 int) (soma, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	soma, subtracao := calculosMatematicos(3, 6)
	fmt.Println(soma)
	fmt.Println(subtracao)
}
