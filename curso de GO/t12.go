package main


import (
	"fmt"
)


//ESTA ES LA FORMA DE CREAR UNA ESTRUCTURA EN GO
//nota : siempre las estructuras se declaran por fuera de la funcion main
type Persona struct{
	nombre string
	apellido string
	edad int
}



//TUTORIAL 12 STRUCTS EN GO O ESTRUCTURAS
func main() {
	// crear un nuevo objeto segun la estructura Persona creada
	persona1 := Persona{nombre:"pablo", apellido:"taddei", edad: 27}
	fmt.Println(persona1)
	
	//segunda forma de crear un objeto con new
	persona2 := new(Persona)
	persona2.nombre = "Lilian"

	fmt.Println(persona2.nombre)
	

}