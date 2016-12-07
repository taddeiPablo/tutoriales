package main

import (
	"fmt"
	"bufio"
	"os"
)

//TUTORIAL 19 MANEJO DE ARCHIVOS  UTILIZANDO OTROS METODOS QUE NOS PERMITEN
//LEER UN ARCHIVO LINEA POR LINEA
func main() {
	//abrimos el archivo
	archivo, error := os.Open("./test.txt")

	if error != nil{
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