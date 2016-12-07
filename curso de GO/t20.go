package main


import (
	"fmt"
	"bufio"
	"os"
)

//TUTORIAL 20 EN ESTE CASO SE APRENDERA EL MANEJO DE DEFER Y PANIC
func main() {
	//abrimos el archivo
	archivo, error := os.Open("./tests.txt")
	//aqui esta la implementacion del defer para cerrar el archivo una vez utilizado
	defer func () {
		archivo.Close()
		fmt.Println("se cerro el archivo")	
		r := recover()
		fmt.Println(r)
	}()

	if error != nil{
		panic(error)
		fmt.Println("Hubo error")
	}
	//hacemos un scanner del archivo
	scanner := bufio.NewScanner(archivo)

	//recorremos el contenido del archivo
	for scanner.Scan(){
		//leemos la info del archivo
		linea := scanner.Text()
		fmt.Println(linea)
	}
}