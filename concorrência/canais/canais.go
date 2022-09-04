package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	canal := make(chan string)

	fmt.Println("Inicio do programa")
	go escrever("Olá Mundo", canal)

	// for {
	// 	texto, aberto := <-canal
	// 	if !aberto {
	// 		break
	// 	}
	// 	fmt.Println(texto)
	// }

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa")
}

func escrever(texto string, canal chan string) {
	fmt.Println("Escrevendo texto")
	time.Sleep(2 * time.Second)
	for i := 0; i < 5; i++ {
		canal <- strconv.Itoa(i)
		time.Sleep(1 * time.Second)
	}

	close(canal)
}
