package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     int16
}

func main() {
	fmt.Println("Arquivo de struct")

	user1 := usuario{"Valter", 22, endereco{"Rua Dos Bobos", 213}}
	fmt.Println(user1)
}
