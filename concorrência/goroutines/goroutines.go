package main

import (
	"fmt"
	"time"
)

func imprimeNaTela(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

func main() {
	fmt.Println("Inicio das rotinas")
	// goroutines são rotinas criadas pelo comando go que executando o comando após eles em CONCORRÊNCIA
	// com o comando com o go na frente

	// CONCORRÊNCIA = não espera o comando terminar para executar o próximo
	go imprimeNaTela("Olá Mundo") // goroutine
	imprimeNaTela("Sambaio")
}
