package main

import "fmt"

func funcao1() {
	fmt.Println("Função 1 chamada")
}

func funcao2() {
	fmt.Println("Função 2 chamada")
}

func alunoPassouNaMateria(n1, n2 float32) bool {

	// esse print será exibido antes do retorno da função ser executado
	// estranho né ???
	defer fmt.Println("O resultado foi exibido")

	if media := (n1 + n2) / 2; media >= 6 {
		return true
	}

	return false
}

func main() {
	fmt.Println("Aprendendo defer na prática")

	// defer é uma clausula para atrasar a execução de uma função
	// geralmente utilizada pra finalizar conexão com banco de dados e coisas do tipo

	defer funcao1()
	funcao2()

	fmt.Println(alunoPassouNaMateria(3.6, 9.0))
}
