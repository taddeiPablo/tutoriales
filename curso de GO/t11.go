package main

import (
	"fmt"
)


//TUTORIAL 11 EN ESTE CASO VEMOS EL TEMA DE PUNTEROS
func main() {
	/*
		1 - El puntero es una direccion de memoria
		2 - En lugar del valor, tenemos la direccion en la que esta el valor
		3 - x,y => as123d => 5
		4 - X as123d => 6 modifica el valor
		5 - Y  => 6 accede al valor

		*T => *int, *string, *struct
		Valor zero => nil
	*/

	//declaracion de los punteros que van a apuntar a una 
	//variable de tipo entero
	var x,y *int

	//variable de tipo entero
	entero := 5

	//aqui se asigna la direccion de memoria al puntero
	//de la variable entero no el valor que almacena.
	x = &entero
	y = &entero 

	fmt.Println(*x)

	fmt.Println(x)
	fmt.Println(y)

	//AHORA VAMOS A CAMBIAR EL VALOR DE X,Y

	//aqui asignamos un nuevo valor a la variable entero
	//apartir del puntero x en este caso
	*x = 6

	fmt.Println(*x)
	fmt.Println(*y)
	fmt.Println("valor dela varible")
	fmt.Println(entero)
}