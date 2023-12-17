package main

import (
	"fmt"
	"net"
)

func main() {
	serverAddress := "127.0.0.1"
	serverPort := 9090

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverAddress, serverPort))
	if err != nil {
		fmt.Printf("error connect to server:%v\n", err)
		return
	}
	defer conn.Close()

	fmt.Println("connected to the server")

	num := 1
	message := make([]byte, num)
	for i := 0; i < num; i++ {
		message[i] = byte('A')
	}

	for i := 0; i < 10000; i++ {
		_, err = conn.Write(message)
		if err != nil {
			fmt.Printf("error send data to server:%v\n", err)
			return
		}

		buffer := make([]byte, 2048)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("error read server response:%v\n", err)
			return
		}

		response := buffer[:n]
		fmt.Printf("received response from server: %s\n", response)
	}
}
