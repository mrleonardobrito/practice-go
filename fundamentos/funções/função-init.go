package main

import "fmt"

// a função init é a função que executa primeiro no arquivo
// utilizada para iniciar variáveis

func init() {
	fmt.Println("Executando a função init")
}

func main() {
	fmt.Println("Executando a função main")
}
