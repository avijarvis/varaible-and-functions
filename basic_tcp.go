package main

// Import statements should be at the top of the file
import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	// Start listening for incoming TCP connections on port 8080
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	fmt.Println("Server is listening on port 8080...")

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		// Handle the connection in a new goroutine
		go func(conn net.Conn) {
			defer conn.Close() // Ensure the connection is closed when we're done

			// Write a message to the client
			_, err := io.WriteString(conn, "\nHello from TCP server\n")
			if err != nil {
				log.Println("Error writing to client:", err)
				return
			}

			// Print a message to the server console
			fmt.Println("Client connected:", conn.RemoteAddr())

			// Send a formatted message to the client
			_, err = fmt.Fprintf(conn, "Well, I hope!\n")
			if err != nil {
				log.Println("Error writing to client:", err)
			}
		}(conn)
	}
}
