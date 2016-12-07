package main

//aqui ejemplo de como importar paquetes en go
import (
	"fmt"
	"./controllers"
)

//TUTORIAL ULTIMO EN EL CUAL APRENDEREMOS A IMPORTAR PAQUETES
//EN GO Y UTILIZANLOS EN UN PROYECTO
func main() {
	fmt.Println(controllers.Mostrar())
	fmt.Println(controllers.Hola())
}