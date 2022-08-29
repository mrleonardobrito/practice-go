package main

import "fmt"

func main() {

	//SEM PONTEIRO
	variavel1 := "eu sou uma variável"
	variavel2 := variavel1

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	variavel1 = "Eu não sou uma variável"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	//COM PONTEIRO
	var v1 string = "v1 declarada"
	var v2 *string = &v1

	fmt.Println(v1)
	fmt.Println(*v2)

	v1 = "v1 mudada"

	fmt.Println(v1)
	fmt.Println(*v2)
}
