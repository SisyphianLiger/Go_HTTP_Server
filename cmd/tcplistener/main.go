package main

import (
	"fmt"
	"log"
	"net"
)


const PORT = ":42069"

func main() {
	

	tcp_connection, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatalf("Error making TCP connection to PORT: %v, Error: %v\n", PORT, err.Error())
	}

	defer tcp_connection.Close()


	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	for {

		conn, err := tcp_connection.Accept()
		if err != nil {
			log.Fatalf("Error Accepting Request: %v", err.Error())
		}
		defer conn.Close()

		if err != nil {
			log.Fatalf("Error accepting Connection: %v\n", err.Error())
			break
		}
		fmt.Println("Accepted Conntection from", conn.RemoteAddr())
		
		printChannels(getLinesChannel(conn))
	
	}
	fmt.Printf("Closing Connection from: %s", PORT)
	



}
