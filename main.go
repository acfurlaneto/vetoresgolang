package main

import "fmt"

func main() {
    numeros := []int{10, 20, 30, 40, 50}
    numeros = append(numeros, 60, 70, 80)
    fmt.Println("Slice de números:", numeros)
    fmt.Println("Tamanho do slice:", len(numeros))
    fmt.Println("Capacidade do slice:", cap(numeros))
}
