package main

import (
	"fmt"
)

// tutorial 17 veremos la forma de comunicar los go-rutimes atravez de 
// channels
func main() {
	//forma de crear un canal (en el make declaramos el  chan y el tipo de dato)
	channel := make(chan string)

	//generamos una go-rutime, pasandole tambien el canal, chan, y tipo de dato
	go func (channel chan string) {
		for{
			var name string
			fmt.Println("ingrese algo a este canal :")
			fmt.Scanln(&name)
			//aqui estamos enviando informacion por el canal
			channel <- name
		}
	}(channel)	//aqui determinamos el canal a usar

	//aqui estamos a la espera de resivir la info del canal
	msg := <- channel

	//aqui mostramos la info resivida por el canal
	fmt.Println("aqui se muestra la ingresado en el canal")
	fmt.Println(msg)
}