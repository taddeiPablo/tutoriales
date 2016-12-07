package main

import (
	"fmt"
	"strconv"
)


// CONVERSION DE DATOS EN GO CUARTO TUTORIAL
func main() {
	//Primer ejemplo de conversion
	edad := 27
	//convertir un entero a string
	edad_str := strconv.Itoa(edad)

	fmt.Println("Mi edad es : "+edad_str)

	//===================================
	//segundo ejemplo de conversion
	nuevaEdad := "27"
	// aqui se convierte un string a entero y se utiiliza un operador
	// especial el _ para manejar un error que en realidad se desecha.
	nuevaEdad_int,_ := strconv.Atoi(nuevaEdad)

	fmt.Println(nuevaEdad_int+10)
}