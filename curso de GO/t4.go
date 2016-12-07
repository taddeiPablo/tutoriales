package main

import(
	"fmt"
	"bufio"
	"os"
)

//CUARTO TUTORIAL EN EL CUAL SE APRENDIO A LEER DATOS
//DESDE LA CONSOLA
func main() {
	//Primer ejemplo de como leer datos desde la consola
	edad := 27
	fmt.Printf("mi edad es : %d\n",edad)
	var nedad int
	fmt.Println("Coloca tu edad: ")
	fmt.Scanf("%d\n",&nedad)
	fmt.Printf("Tu edad es %d\n", nedad)
	//Segundo ejemplo de como leer datos desde la consola
	// haciendo un reader del paquete bufio

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingrese su nombre")
	text,err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Hola "+nombre )
	}

}