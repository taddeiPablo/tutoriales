package controllers

//importando de librerias necesarias
import (
	"net/http"
	"fmt"
)

//MANEJADOR DE URL /
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ingreso aqui")
	http.ServeFile(w,r,"index.html")
	//otra forma de poder obtener los archivos en go
	// r.URL.Path[1:] ESTO TOMA LA URL DEL REQUEST DESPUES DE LA /
}