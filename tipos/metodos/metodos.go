package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeComleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	p1 := pessoa{
		nome:      "Jeferson",
		sobrenome: "Ferreira",
	}
	fmt.Println(p1.getNomeComleto())

	p1.setNomeCompleto("Jeferson Farias")
	fmt.Println(p1.getNomeComleto())

}
