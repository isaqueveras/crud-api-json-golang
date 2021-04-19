package main

import (
	"fmt"
	"net/http"
)

func rotaPrincipal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Olá mundo!")
}

func configurarRotas() {
	http.HandleFunc("/", rotaPrincipal)
}

func configurarServidor() {
	configurarRotas()

	fmt.Println("Servidor rodando na porta 1317")
	http.ListenAndServe(":1337", nil)
}

func main() {
	configurarServidor()
}
