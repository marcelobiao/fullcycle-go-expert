package main

import "fmt"

func main() {
	salarios := map[string]int{
		"marcelo": 1000,
		"joão":    2000,
		"maria":   3000,
	}

	fmt.Println(salarios["marcelo"])
	delete(salarios, "marcelo")
	salarios["pedro"] = 1500
	fmt.Println(salarios["pedro"])

	// Maneiras de inicializar um map
	// sal := make(map[string]int)
	// sal1 := map[string]int{}

	for nome, valor := range salarios {
		fmt.Printf("Nome: %s valor: %d \n", nome, valor)
	}
}
