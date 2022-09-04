package main

import (
	"fmt"
	"time"
)

func main() {
	canal := multiplexar(escrever("Multiplexador 1"), escrever("Multiplexador 2"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func multiplexar(canal1, canal2 <-chan string) <-chan string {
	canalSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canal1:
				canalSaida <- mensagem
			case mensagem := <-canal2:
				canalSaida <- mensagem
			}
		}
	}()

	return canalSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Second / 2)
		}
	}()

	return canal
}
