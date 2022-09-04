package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return 1
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func worker(tarefas <-chan uint, resultados chan<- uint) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func main() {
	quant := 45
	tarefas, resultados := make(chan uint, 45), make(chan uint, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < quant; i++ {
		tarefas <- uint(i)
	}
	close(tarefas)

	for i := 0; i < quant; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}
