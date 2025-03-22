package main

import (
	"fmt"
	"log"
	"net"
)


const PORT = ":42069"

func main() {
	

	tcp_connection, err := net.Listen("tcp", PORT)
	defer tcp_connection.Close()

	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	for {
		conn, err := tcp_connection.Accept()
		if err != nil {
			log.Fatalf("Error accepting Connection: %v\n", err.Error())
			break
		}
		fmt.Println("Accepted Conntection from", conn.RemoteAddr())
		
		printChannels(getLinesChannel(conn))
	
	}
	fmt.Printf("Closing Connection from: %s", PORT)
	



}
