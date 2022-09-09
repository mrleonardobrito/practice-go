package main

import (
	"fmt"
	"testes-automatizados/enderecos"
)

func main() {
	tipoDoEnderecoValido := enderecos.TipoDeEndereco("Rua Imigrantes, 331")

	fmt.Println(tipoDoEnderecoValido)
}
