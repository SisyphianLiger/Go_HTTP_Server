package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)



func main() {
	
	serverAddr := "localhost:42069"

	local_address, err := net.ResolveUDPAddr("udp", serverAddr)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error Resolving UDP address: %v\n", err)
		os.Exit(1)
	}

		
	udpConn, err := net.DialUDP("udp", nil, local_address)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error dialing UDP address: %v\n", err)
		os.Exit(1)
	}
	
	defer udpConn.Close()
	

	
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("> ")
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			os.Exit(1)
		}

		_, err_conn := udpConn.Write([]byte(input))
		if err_conn != nil {
			fmt.Fprintf(os.Stderr, "Error sending message: %v\n", err_conn)
			os.Exit(1)
		}

		fmt.Printf("Message sent: %s", input)


	}
	
}
