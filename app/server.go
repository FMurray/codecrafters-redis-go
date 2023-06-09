package main

import (
	"fmt"
	// Uncomment this block to pass the first stage
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	buff := make([]byte, 4096)
	tmp := make([]byte, 1024)
	for {
		n, err := c.Read(tmp)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			os.Exit(1)
		}
		fmt.Println("Read:", n)
		const Ok = "+PONG\r\n"
		c.Write([]byte(Ok))
		buff = append(buff, tmp[:n]...)
	}
	fmt.Println("Received:", string(buff))


	c.Close()
}