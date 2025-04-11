package main

import "fmt"

var saldo int

func saudacao(nome string) {
fmt.Println("\nOlá,", nome, "👋")
fmt.Println("Bem-vindo(a) ao seu banco digital!")
}
func sacar(valor int) {
if valor > saldo {
fmt.Println("Saldo insuficiente!")
} else {
saldo -= valor
fmt.Println("Saque realizado com sucesso!")
	}
}
func depositar(valor int) {
saldo += valor
fmt.Println("Depósito realizado com sucesso!")
}
func verSaldo() {
fmt.Println("Saldo atual:", saldo)
}
func mostrarMenu() {
fmt.Println("O que você deseja fazer?")
fmt.Println("1 - Sacar")
fmt.Println("2 - Depositar")
fmt.Println("3 - Ver saldo")
fmt.Println("4 - Sair")
fmt.Print("Escolha uma opção: ")
}
func main() {
var nome string
var opcao int
var valor int

	fmt.Print("Digite seu nome: ")
	fmt.Scan(&nome)
	saudacao(nome)
	fmt.Print("\nDigite seu saldo inicial: ")
	fmt.Scan(&saldo)

	for {

	mostrarMenu()
	fmt.Scan(&opcao)
switch opcao {
case 1:
fmt.Print("Digite o valor para sacar: ")
fmt.Scan(&valor)
sacar(valor)
case 2:
fmt.Print("Digite o valor para depositar: ")
fmt.Scan(&valor)
depositar(valor)
case 3:
verSaldo()
case 4:
fmt.Println("Até logo,", nome+"!")
return
default:
fmt.Println(" Opção inválida! Tente novamente.")
	}
}
}
