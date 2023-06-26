package main

import "fmt"

type user struct {
	nome     string
	idade    uint8
	sexo     bool // true masculino, false feminino
	endereco localizacao
}

type localizacao struct {
	rua    string
	numero uint16
	bairro string
	cidade string
	estado string
	CEP    uint32
}

func main() {
	fmt.Println("Stored Data Management")

	endereco_usuario := localizacao{rua: "Rua tananana",
		numero: 99,
		bairro: "Tal Bairro",
		cidade: "Niterói",
		estado: "RJ",
		CEP:    12345000}

	usuario := user{nome: "Lucas Lourenço",
		idade:    20,
		sexo:     true,
		endereco: endereco_usuario}

	fmt.Println(usuario)
}
