package main

import "fmt"

func main() {
	canal := make(chan string, 3)

	canal <- "Olá mundo"
	canal <- "Canal com buffer de 3 espaços"
	canal <- "Não dá deadlock"

	mensagem := <-canal
	mensagem2 := <-canal
	mensagem3 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
	fmt.Println(mensagem3)
}
