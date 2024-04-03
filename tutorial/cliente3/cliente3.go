package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

type Persona struct {
	Nombre string
	Email  []string
}

func cliente(persona Persona) {
	c, err := net.Dial("tcp", ":9999")

	if err != nil {
		fmt.Println(err)
		return
	}
	err = gob.NewEncoder(c).Encode(persona)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func main() {
	persona := Persona{
		Nombre: "Eduar",
		Email: []string{
			"corre1@hotmail.com",
			"correo2@hotmail.com",
		},
	}

	go cliente(persona)

	var input string
	fmt.Scanln(&input)
}
