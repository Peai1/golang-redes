package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func cliente() {
	c, err := net.Dial("tcp", ":9999")

	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Open("img.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = io.Copy(c, file) // copia en c lo que recibe de file
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Imagen enviada!")

	c.Close()
}

func main() {
	go cliente()

	var input string
	fmt.Scanln(&input)
}
