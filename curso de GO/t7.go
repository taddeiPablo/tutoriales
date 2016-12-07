package main

import (
	"fmt"
)

// SEPTIMO TUTORIAL EN ESTE CASO ARREGLOS ARRAYS
func main() {
	//declaracion de un array vacio
	var arreglo [3]int
	//declaracion de un array inicializado con valores
	arreglo1 := [3]int{1,2,3}
	
	//declaracion de un array multidimencional o matriz
	var matriz [4][2]int

	fmt.Println(arreglo)
	fmt.Println(arreglo1)
	fmt.Println(len(arreglo))	
	fmt.Println(matriz)
}