package main

import (
	"fmt"
)

// TUTORIAL 10 EN ESTE CASO VEMOS LA COPIA DE SLICE
func main() {
	/*
		Copy : copia el minimo de elementos en ambos arreglos
	*/

	slice := []int{1,2,3,4}
	copia := make([]int, len(slice))
	copy(copia, slice)

	fmt.Println(slice)
	fmt.Println(copia)
}