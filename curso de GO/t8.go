package main

import (
	"fmt"
)

//declaracion de slice en go
func main() {
	matriz := []int{1,2}
	if matriz == nil{
		fmt.Println("esta vacio")
	}else{
		fmt.Println(len(matriz))
	}
	//puntero al arreglo
	//longitud del arreglo al que apunta
	//capacidad
	//----------------------------------
	//ahora ejemplo de slice
	var arreglo [3]int
	//apartir de esta operacion partimos un array y nos traemos los datos desde la posicion que determinemos
	slice := arreglo[:2] 
	fmt.Println(slice)

}