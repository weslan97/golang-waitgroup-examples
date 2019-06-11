package main

import (
	"fmt"
	"sync"
)

func facaAlgo(w *sync.WaitGroup) {
	fmt.Println("   Feito!")
	w.Done()
}

func main() {
	// OBS: Executar em um ambiente com vários núcleos de processamento
	var w sync.WaitGroup
	fmt.Println("Inicio")
	w.Add(1)
	go facaAlgo(&w)
	w.Wait()
	fmt.Println("Fim")
}
