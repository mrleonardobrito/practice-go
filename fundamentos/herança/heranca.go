package main

import "fmt"

type pessoa struct {
	nome   string
	idade  uint8
	altura uint16
}

type estudante struct {
	pessoa
	faculdade string
	curso     string
	periodo   int8
}

func main() {
	pessoaExemplo := pessoa{"Rafaela Nunes", 18, 168}
	estudanteExemplo := estudante{pessoaExemplo, "USP", "jornalismo", 6}

	fmt.Println(estudanteExemplo)
	fmt.Println(estudanteExemplo.nome)
	fmt.Println(estudanteExemplo.idade)
}
