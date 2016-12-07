package main

import (
	"net/http"
	"io"
)

//TUTORIAL 21 VEMOS COMO ARMAR UN SERVIDOR WEB MUY SIMPLE
// COMO IR ARMANDO LAS ROUTES..
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8000", nil)
}

//PRIMER MANEJADOR RESPONDE A /
func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>HOLA MUNDO</h1>")
}

//segundo manejador responde a /login
func login(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "en breve ya va a estar funcionando la pagina de logueo")
}