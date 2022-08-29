package main

import (
	"errors"
	"fmt"
)

func main() {
	// INTEIROS
	var numero1 int8 = 100
	var numero2 int16 = 10000
	var numero3 int32 = 1000000000
	var numero4 int64 = 1000000000000000000

	fmt.Println(numero1, numero2, numero3, numero4)

	var numeroint int = 10000000000 // o int usa a arquitetura do computador ex: x64 será um int64

	fmt.Println(numeroint)

	//unsigned ints

	var unumero1 uint8 = 100
	var unumero2 uint16 = 10000
	var unumero3 uint32 = 1000000000
	var unumero4 uint64 = 1000000000000000000

	fmt.Println(unumero1, unumero2, unumero3, unumero4)

	//aliases para inteiros

	var intrune rune = 100000000 //aliase para int32
	var intbyte byte = 100       //aliase para uint8

	fmt.Println(intrune, intbyte)

	//REAIS

	//float32, float64
	var real1 float32 = 100.00
	var real2 float64 = 1000000.000
	real3 := 12311.123123

	fmt.Println(real1, real2, real3)

	// STRINGS

	var str string = "String1"
	str2 := "String2"
	char := 'B'

	fmt.Println(str, str2, char)

	// BOOLEAN

	var booleano1 bool = true
	fmt.Println(booleano1)

	//ERROR

	var erro error = errors.New("Erro Interno")
	fmt.Println(erro)
}
