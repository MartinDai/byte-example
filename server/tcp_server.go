package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	address := "0.0.0.0"
	port := 9090
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		fmt.Printf("error create listener:%v\n", err)
		return
	}

	defer listener.Close()

	fmt.Printf("server is listening on %s:%d\n", address, port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("error accept connection:%v\n", err)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 2048)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("client closed the connection")
				return
			}
			fmt.Printf("error read from client: %v\n", err)
			return
		}

		data := buffer[:n]
		fmt.Printf("received data:%s\n", data)

		_, err = conn.Write(data)
		if err != nil {
			fmt.Printf("error send response data to client:%v\n", err)
			return
		}
	}
}
