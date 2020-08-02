package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$%2.f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {

	var coisa imprimivel = pessoa{"Roberto", "Silva"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{"Calça jean", 79.90}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	pessoa := pessoa{
		nome:      "Jeferson",
		sobrenome: "Ferreira",
	}

	produto := produto{
		nome:  "Lapis",
		preco: 3.5,
	}

	imprimir(pessoa)
	imprimir(produto)

}
