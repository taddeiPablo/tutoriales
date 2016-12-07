package main


import(
	"fmt"
)

// aqui la declaracion de un tipo que va a servir de tipo base para la herencia hacia otros tipos
type Humano struct{
	name string
}

// aqui se declara una funcion que sera heredada por los tipos que deriven de humano
func (this Humano) hablar() string{
	return "este es un humano"
}


// aqui se declara el tipo que hereda de humano, implementara los atributos y metodos de humano
type Persona struct{
	Humano
}

// aqui la forma de sobre-escribir un metodo en go este metodo es de humano
func (this Persona) hablar() string {
	this.Humano.hablar()
	return "Bienvenidos a este simple ejemplo de sobre escritura"
}

// TUTORIAL NUMERO 14 ESTAMOS VIENDO CAMPOS ANONIMOS
//ademas se va a estar implementado herencia
func main() {
	//
	tutor := Persona{Humano{"pablo"}}

	fmt.Println(tutor.Humano.name)
	fmt.Println(tutor.hablar())

}