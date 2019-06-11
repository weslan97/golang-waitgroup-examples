package main

import (
	"fmt"
)

func facaAlgo() {
	fmt.Println("   Feito!")
}

func main() {
	// OBS: Executar em um ambiente com vários núcleos de processamento
	fmt.Println("Inicio")
	go facaAlgo()
	fmt.Println("Fim")
}
