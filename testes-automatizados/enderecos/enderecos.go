package enderecos

import "strings"

// TipoDeEndereco retorna a primeira palavra de um endereço válido
func TipoDeEndereco(endereco string) string {
	primeiraPalavraDoEndereco := strings.Split(endereco, " ")[0]
	sliceValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	for _, tipoValido := range sliceValidos {
		if palavra := strings.ToLower(primeiraPalavraDoEndereco); palavra == tipoValido {
			return strings.Title(palavra)
		}
	}

	return "Tipo Inválido"
}
