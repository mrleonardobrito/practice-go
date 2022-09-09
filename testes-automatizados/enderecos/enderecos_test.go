package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoEnviado  string
	enderecoEsperado string
}

func TestTipoDeEndereco(t *testing.T) {
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua Walter Bezerra", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Escola Pedro Bezerra", "Tipo Inválido"},
		{"RUA WALTER BEZERRA", "Rua"},
		{"ESTRADA MULIRO COUTO", "Estrada"},
		{"ROdOviA passa Anerl", "Rodovia"},
		{"", "Tipo Inválido"},
		{"1231", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		resultadoRecebido := TipoDeEndereco(cenario.enderecoEnviado)
		if resultadoRecebido != cenario.enderecoEsperado {
			t.Errorf("Endereço enviado: %s Resutado esperado: %s Resultado recebido: %s \n",
				cenario.enderecoEnviado,
				cenario.enderecoEsperado,
				resultadoRecebido)
		}
	}
}
