package main

import "fmt"

func main() {
	var saldo, valor float64
	var acao string

	fmt.Print("Saldo inicial: R$ ")
	fmt.Scan(&saldo)

	for {
fmt.Print("\nDigite 'sacar', 'depositar' ou 'sair': ")
fmt.Scan(&acao)

if acao == "sair" {
break
}

fmt.Print("Valor: R$ ")
fmt.Scan(&valor)

if acao == "depositar" {
saldo += valor
} else if acao == "sacar" {
if valor <= saldo {
saldo -= valor
} else {
fmt.Println("Saldo insuficiente.")
}
} else {
fmt.Println("Ação inválida.")
}
fmt.Printf("Saldo atual: R$ %.2f\n", saldo)
	}
}
