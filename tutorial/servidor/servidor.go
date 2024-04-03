package main

import (
	"fmt"
	"net"
)

func servidor() {
	s, err := net.Listen("tcp", ":9999") // s -> servidor
	// en este puerto se va a escuchar las peticiones, por cada cliente se crea un hilo que ejecuta la funcion handleClient

	if err != nil {
		fmt.Println(err)
		return
	}

	for { // Loop infinito para aceptar conexiones
		c, err := s.Accept() // c -> cliente
		if err != nil {
			fmt.Println(err)
			continue // Si hay un error, se sigue con el loop
		}
		go handleClient(c) // por cada cliente, se crea un hilo
		// el go genera una hebrea que ejecuta la funcion handleClient
	}

}

func handleClient(c net.Conn) {
	b := make([]byte, 100) // buffer de 100 bytes, lee hasta 100 bytes
	bs, err := c.Read(b)   // lee el mensaje del cliente y lo guarda en b

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Mensaje: ", string(b[:bs])) // imprime el mensaje del cliente, desde posicion 0 hasta cantidad de bytes leidos, sino imprime basura
	fmt.Println("Bytes: ", bs)               // imprime la cantidad de bytes leidos
}

func main() {
	go servidor()

	var input string
	fmt.Scanln(&input)

}
