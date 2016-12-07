package main

import (
	"fmt"
)

//Aqui se crear un tipo Persona
type Persona struct{
	nombre string
	apellido string
	edad int
}

//Aqui se crea una funcion que se asigna al tipo persona
func (this Persona) nombre_completo() string{
	return this.nombre + " " + this.apellido
}

//Aqui se crear una funcion setter para modificar la variable nombre
func (this *Persona) set_name(n string){
	this.nombre = n
}

//TUTORIAL 13 EN ESTE CASO VEREMOS COMO CREAR METODOS PARA LAS STRUCT
func main() {
	persona1 := new(Persona)

	persona1.nombre = "pablo"
	persona1.apellido = "taddei"

	//asi llamamos a la funcion setter
	persona1.set_name("Pablo Alejandro")

	fmt.Println(persona1.nombre_completo())
}