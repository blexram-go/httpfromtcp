package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const port = ":42069"

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		log.Fatalf("Error resolving UDP address: %v", err)
	}

	udpConn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Fatalf("Error dialing UDP connection: %v", err)
	}
	defer udpConn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		userInput, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error reading user input: %v", err)
		}
		if _, err := udpConn.Write([]byte(userInput)); err != nil {
			log.Fatalf("Error writing to UDP connection: %v", err)
		}
		fmt.Printf("Message sent: %s", userInput)
	}
}
