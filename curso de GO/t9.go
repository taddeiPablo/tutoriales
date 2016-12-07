package main

import (
	"fmt"
)


//utilizacion de make y append utlizando slice
func main() {
	arreglo := [3]int{1,2,3}
	slice := arreglo[0:]
	fmt.Println(slice)

	//aqui un ejemplo de como utilizar un make
	slice1 := make([]int, 3, 5) //como tercer parametro el capasity .
	fmt.Println(cap(slice1))

	slice1 = append(slice1,2)
	fmt.Println(slice1)

}