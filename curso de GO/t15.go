package main


import (
	"fmt"
)

// de esta manera se definen interfaces en go
type Iuser interface{
	Permisos() int //nivel de permisos del 1 - 5
	Nombre() string
}

//aqui definicion de una estructura en go
type Admin struct{
	name string
}

//aqui este metodo es el implementado de la interface
// implementamos este metodo de la interfaces en la estructura admin
func (this Admin) Permisos() int {
	return 5
}

//aqui este metodo es el implementado de la interface
// implementamos este metodo de la interfaces en la estructura admin
func (this Admin) Nombre() string {
	return this.name
}


type Editor struct{
	name string
}


//aqui este metodo es el implementado de la interface
// implementamos este metodo de la interfaces en la estructura editor
func (this Editor) Permisos() int {
	return 3
}

//aqui este metodo es el implementado de la interface
// implementamos este metodo de la interfaces en la estructura editor
func (this Editor) Nombre() string {
	return this.name
}





//funcion que no pertenese a ninguna interfaces o estrutura
func auth(user Iuser) string {
	if user.Permisos() == 5{
		return user.Nombre() + "  Tiene permisos de admin"
	}else if user.Permisos() == 3{
		return user.Nombre() + "  posse pemisos limitados este usuario es un editor"
	}else{
		return "no posee ningun permiso"
	}
}

//TUTORIAL NUMERO 15 MANEJO DE INTERFACES
func main() {
	//ejemplo de como utilizar las estructuras que implementan la
	//interface
	admin := Admin{"pablo"}
	editor := Editor{"lala"}
	fmt.Println(auth(admin))
	fmt.Println(auth(editor))
	//segundo ejemplo con array
	//aqui se puede ver como crear un array del tipo de la interces
	// y finalmente la inicializamos con dos estructuras que implementan
	// esa interfaces
	usuarios := []Iuser{Admin{"pablo"},Editor{"lala"}}

	//ejemplo de for con el cual iteramos el array
	for _,usuario := range usuarios{
		fmt.Println(auth(usuario))
	}
}

