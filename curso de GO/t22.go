package main


import (
	"net/http"
	"log"
	"encoding/json"
)


type Curso struct{
	Nombre string `json:"nombre"`
	Version int
}

type Cursos []Curso


func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/login", login)
	http.HandleFunc("/get", query)
	http.HandleFunc("/datos", test)


	log.Println("Listening...")
	log.Println("Port : 8001")
	//forma de configurar a go para que solo responda con archivos 
	//publicos
	fileServer := http.FileServer(http.Dir("public"))
	http.Handle("/public/",http.StripPrefix("/public/",fileServer))
	log.Println(http.ListenAndServe(":8001", nil))
}


func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r,"index.html")
	//otra forma de poder obtener los archivos en go
	// r.URL.Path[1:] ESTO TOMA LA URL DEL REQUEST DESPUES DE LA /
}

func login(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r,"login.html")
}

func query(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":	
		// server
		http.ServeFile(w,r,"post.html")
		log.Println("get")
		log.Println(r.URL.Path)
	case "POST":
		//server
		
		log.Println("POST : /json")
	case "PUT":
		//server
	case "DELETE":
		//server
	}
}

func test(w http.ResponseWriter, r*http.Request) {
	switch r.Method{
		case "GET":
			curso1 := Curso{"programacion en GOLANG", 1}
			json.NewEncoder(w).Encode(curso1)
		case "POST":
			log.Println("post: ingreso por aqui")
			curso2 := Curso{"programacion en GOLANG", 1}
			json.NewEncoder(w).Encode(curso2)
	}
}