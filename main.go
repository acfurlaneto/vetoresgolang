package main

import "fmt"

func main() {
    nomes := []string{"Ana", "Paola", "Rubi", "Teresa", "Clark"}
    fmt.Println (nomes)
    rangeOne := nomes[:2]
    fmt.Println(rangeOne) 
    rangeTwo := nomes[3:]
    fmt.Println(rangeTwo) 
    rangeThree := nomes[2:3] 
    fmt.Println(rangeThree) 
 
}

