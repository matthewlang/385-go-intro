package main

import (
	"io"
	"log"
	"net"
	"time"
)

// Handler for socket connections.
func handleConn(c net.Conn) {
	defer c.Close()

	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			log.Printf("%v disconnected", c.RemoteAddr())
			return // client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// Listen for connections on port 8000.
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Error listening on 8000: %v", err)
	}

	// Continuously accept and handle connections.
	for {
		// Wait for and accept a connection on the listener.
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error: %v", err)
			continue
		}
		log.Printf("Connection from %v", conn.RemoteAddr())
		// Dispatch handling the connection the handler.
		go handleConn(conn)
		log.Printf("Listening again...\n")
	}
}
