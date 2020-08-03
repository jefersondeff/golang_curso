package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct para json
	p1 := produto{1, "Notebook", 1899.90, []string{"Promoção", "Eletrônico"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	// json para struct
	var p2 produto
	jsonString := `{"id": 2, "nome": "caneta", "preco": 8.90, "tags": ["Papelaria", "Importado"]}`
	err := json.Unmarshal([]byte(jsonString), &p2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(p2)
}
