package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

func imprimeArea(f forma) {
	fmt.Println("A area da forma é ", f.area())
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func main() {
	c1 := circulo{10.0}
	r1 := retangulo{10.0, 20.0}

	imprimeArea(c1)
	imprimeArea(r1)
}
