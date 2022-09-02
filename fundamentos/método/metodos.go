package main

import "fmt"

// métodos são funções vinculadas a tipos
type usuario struct {
	nome  string
	idade uint8
}

// o método está vinculado ao tipo usuário e pode ser chamado pela instância dele
func (u usuario) salvar() {
	fmt.Println("Salvando no banco de dados")
}

// os métodos podem fazer retornos(como as variáveis)
func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

// os métodos podem utilizar o ponteiro da instancia do struct para mudar o valor
func (u *usuario) fazerAniversario() {
	// utilizando o ponteiro do tipo usuario podemos pular o processo de desreferenciação
	u.idade++
}

func main() {
	user1 := usuario{"Rebecca Sugar", 20}
	user2 := usuario{"Raphael Draccon", 17}

	user1.salvar()
	antesDoAniversario := user2.maiorDeIdade()
	fmt.Println(antesDoAniversario)

	user2.fazerAniversario()
	depoisDoAniversario := user2.maiorDeIdade()
	fmt.Println(depoisDoAniversario)

	fmt.Println(user1)
	fmt.Println(user2)
}
