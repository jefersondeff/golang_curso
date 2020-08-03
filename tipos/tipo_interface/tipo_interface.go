package main

import "fmt"

type curso struct {
	name string
}

func main() {
	var coisa interface{}
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)

	type dimanico interface{}
	var coisa2 dimanico = "Opa"
	fmt.Println(coisa2)

	coisa2 = true
	fmt.Println(coisa2)

	coisa2 = curso{"Golang: Explorando a linguagem do google."}
	fmt.Println(coisa2)
}
