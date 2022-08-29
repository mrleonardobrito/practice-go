package main

import "fmt"

func main() {

	usuario1 := map[string]string{
		"nome":      "Marcela",
		"idade":     "21",
		"faculdade": "jornalismo",
	}

	usuario2 := map[string]map[string]string{
		"pessoal": {
			"nome":  "Rebeca Sugar",
			"idade": "21",
		},
		"endereco": {
			"rua":    "chinatown",
			"numero": "21",
		},
	}

	// adicionando informações
	usuario2["signo"] = map[string]string{
		"rua": "Leão",
	}

	// mudando informações
	usuario1["idade"] = "34"

	fmt.Println(usuario1)
	fmt.Println(usuario2)
}
