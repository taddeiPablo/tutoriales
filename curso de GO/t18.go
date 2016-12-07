package main


import (
	"io/ioutil"
	"fmt"
)

//TUTORIAL NUMERO 18 MANEJO DE ARCHIVOS EN GO
func main() {
	file_data,err := ioutil.ReadFile("./test.txt")

	if err != nil{
		fmt.Println("hubo un error")
	}
	fmt.Println(string(file_data))
}