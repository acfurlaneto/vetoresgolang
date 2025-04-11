package main

import "fmt"

func sayGreeting(nome string){
    fmt.Println("Olá!", nome)
}
func addNumber(numero1 int, numero2 int)int{
    return numero1 + numero2

    func versaldo(){
        fmt.Println("seu saldo é:")
    }
    func depositar(){
        fmt.Println("deposite um valor:")
    }
    func sacar(){
        fmt.Println("saque um valor:")
    }
    
    fmt.Print("\nDigite "sacar", "depositar", "ver saldo" ou "sair": ")
    fmt.Scan(&acao)
    
    }

func main() {
    sayGreeting("Juvelino")
    resultado := addNumber(10, 20)
    fmt.Println(resultado)
}
