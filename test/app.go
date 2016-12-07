package main


import (
	"net/http"
	"./controllers"
)


func main() {
	
	/*AQUI EN ESTA SECCION SE DECLARAN LOS HANDLERS URL*/	
	http.HandleFunc("/", controllers.Handler)



	/*AQUI CONFIGURACION DEL SERVER*/
	fileServer := http.FileServer(http.Dir("public"))
	http.Handle("/public/",http.StripPrefix("/public/",fileServer))
	http.ListenAndServe(":8000", nil)
}