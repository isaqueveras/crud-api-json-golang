package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Livro struct {
	Id int
	Titulo string
	Autor string
}

// Criando uma lista
var Livros []Livro = []Livro {
	Livro {
		Id: 1,
		Titulo: "O Guarani",
		Autor: "José de Alencar",
	},
	Livro {
		Id: 2,
		Titulo: "Ponto de inflexão",
		Autor: "Flávio Augusto",
	},
}

func rotaPrincipal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Olá mundo!")
}

func listarLivros(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.Encode(Livros)	
}

func configurarRotas() {
	http.HandleFunc("/", rotaPrincipal)
	http.HandleFunc("/livros", listarLivros)
}

func configurarServidor() {
	configurarRotas()

	fmt.Println("Servidor rodando na porta 1317")
	http.ListenAndServe(":1337", nil)
}

func main() {
	configurarServidor()
}
