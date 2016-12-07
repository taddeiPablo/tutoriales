package main

import (
	"fmt"
	"time"
	"strings"
)

//TUTORIAL GOROUTINES 
//PROGRAMACION CONCURRENTE
func main() {
	//de esta manera se utilizan las GO-Rutines
	go mi_nombre_lento("PABLO")
	fmt.Println("ingrese")
	var wait string
	fmt.Scanln(&wait)
}
	

//funcion que se bloque en un segund
func mi_nombre_lento(nombre string) {
	letras := strings.Split(nombre,"")

	for _,letra := range(letras){
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}