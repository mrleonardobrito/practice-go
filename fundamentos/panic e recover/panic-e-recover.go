package main

import (
	"fmt"
	"time"
)

func iniciaMontanhaRussa() {
	fmt.Println("Montanha Russa iniciada")
}

func verificaSegurança(quantFalhas int) string {
	defer verificaSaudePassageiros(quantFalhas)

	if quantFalhas > 0 {
		panic("HOUVE UMA FALHA NA SEGURANÇA PARANDO MONTANHA RUSSA")
	}

	return "percurso feito com segurança"
}

func verificaSaudePassageiros(quantFalhas int) {
	if r := recover(); r != nil {
		falhas := quantFalhas
		for i := 0; i < quantFalhas; i++ {
			falhas -= i
			fmt.Println(i+1, "Falha(s) resolvida(s)")
			time.Sleep(1 * time.Second)
		}
		fmt.Println("Falhas resolvidas com sucesso, retomando montanha russa!!!")
	}
}

func main() {
	iniciaMontanhaRussa()
	fmt.Println(verificaSegurança(3))
	fmt.Println("Carrinho chegou no final do percurso com segurança!!!")
}
