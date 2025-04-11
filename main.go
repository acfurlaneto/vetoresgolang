package main

import "fmt"

func main() {
	var idade int

	fmt.Print("Informe sua idade: ")
	fmt.Scan(&idade)

if idade < 18 {
fmt.Println("Você ainda é menor de idade 🎀 Aproveita a fase!")
} else if idade <= 60 {
fmt.Println("Você é um adulto! 💼 Bora pagar boletos!.")
} else {
fmt.Println("Você é Idoso!👑👵🏼.")
	}
}
