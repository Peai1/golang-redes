package main

import (
	"fmt"
	"net"
	"strings"
)

func mensaje_bienvenida() {
	fmt.Println("Iniciando servidor")
}

func main() {
	// Emito mensaje
	go mensaje_bienvenida()
	// Veo si puerto no tiene problemas
	PORT := ":63420"
	BUFFER := 1024

	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer connection.Close()
	buffer := make([]byte, BUFFER)

	// While para quedarse escuchando en el puerto, mostrar lo recibido y responder con otro mensaje
	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		fmt.Println("->", string(buffer[0:n]))
		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Cerrando server Golang")
			return
		}

		mensaje := []byte("Golang te saluda...")
		fmt.Printf("data: %s\n", string(mensaje))
		_, err = connection.WriteToUDP(mensaje, addr)
		if err != nil {
			fmt.Println(err)
			return
		}

	}

}
