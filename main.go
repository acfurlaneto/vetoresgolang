package main

import "fmt" 

func main() { 
	estoque := make(map[string]int) 
	estoque["Coxinha"] = 20
	estoque["Pão de Queijo"] = 10
estoque["Refrigerante"] = 20

		estoque ["Coxinha"] -= 2
		estoque ["Pão de Queijo"] -=1
	
	for produto, total := range estoque {
	fmt.Printf("%s : Produto com %d de estoque\n", produto, total) 
	}
	} 

	