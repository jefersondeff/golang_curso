package main

import "fmt"

type produto struct {
	name     string
	preco    float64
	desconto float64
}

//Método: Função  com reiceiver (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {

	var produto1 produto
	produto1 = produto{
		name:     "Lapis",
		preco:    1.79,
		desconto: 0.05,
	}
	fmt.Println(produto1, produto1.precoComDesconto())

	produto2 := produto{
		name:     "Caneta",
		preco:    2.33,
		desconto: 0.3,
	}

	fmt.Println(produto2, produto2.precoComDesconto())
}
