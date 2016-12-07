package main

import (
	"net/http"
	"encoding/json"
	"fmt"
)


type Curso struct{
	Nombre string `json:"nombre"`
	Version int
}

type Cursos []Curso


//TUTORIAL 23 VAMOS A VER EL MANEJO DE JSON EN GO
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		curso1 := Curso{"programacion en GOLANG", 1}
		fmt.Println("ingreso aqui desde la url")
		fmt.Println(curso1)
		json.NewEncoder(w).Encode(curso1)
	})

	http.ListenAndServe(":8000", nil)
}