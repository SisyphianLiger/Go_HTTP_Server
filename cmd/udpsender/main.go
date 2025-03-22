package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)


const PORT = ":42069"

func main() {
	local_address, err := net.ResolveUDPAddr("localhost", ":42069")

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

		
	udpConn, err := net.DialUDP("udp", local_address, nil)
	
	if err != nil {
		log.Fatalf("UDP could not be started: %v", err)
	}

	defer udpConn.Close()
	
	reader := bufio.NewReader(udpConn)
	for {
		fmt.Printf(">\n")
		input, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
		}

		_, error := udpConn.Write([]byte(input))
		if error != nil {
			log.Printf("Error Writing Message: %v", error.Error())
		}


	}
	
}
