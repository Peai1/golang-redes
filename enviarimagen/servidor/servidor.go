package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func servidor() {
	s, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleClient(c)
	}
}

func handleClient(c net.Conn) {
	defer c.Close()

	file, err := os.OpenFile("imagen_recibida.png", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, c) // copia en file lo que recibe de c
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Imagen recibida!")
}

func main() {
	go servidor()

	// var input string
	// fmt.Scanln(&input)
}
