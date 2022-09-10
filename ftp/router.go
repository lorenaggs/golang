// Package ftp provides structs and functions for running an FTP server.
package ftp

//is our application’s router. It’s the workhorse that checks for incoming commands
//and matches them against the FTP routines that the server implements.
import (
	"bufio"
	"log"
	"strings"
)

const (
	MaxBufferMb   = 10
	MaxBufferByte = 10 * 1024 * 1024
)

// Serve scans incoming requests for valid commands and routes them to handler functions.
func Router(conn *Conn) {
	conn.respond(status220) //The first thing we do upon entering Serve is to issue a 220 response to the client,
	conn.printChannels()

	inputClient := bufio.NewScanner(conn.conn) //To listen for incoming commands,
	buffer := make([]byte, MaxBufferByte)
	inputClient.Buffer(buffer, MaxBufferByte)

	for inputClient.Scan() {
		input := strings.Fields(inputClient.Text())
		if len(input) == 0 {
			continue
		}
		if !conn.hasUserChannel() {
			continue
		}
		command, args := input[0], input[1:]  // you can see exactly what the client is sending
		log.Printf("<< %s %v", command, args) // you can see exactly what the client is sending

		switch command {
		case "join":
			conn.joinChannel(args)
		case "list":
			conn.list(args)
		case "port":
			conn.port(args)
		case "user":
			conn.user(args)
		case "exit":
			conn.respond(status221)
			return
		case "send": // get //the client secretly sends a port
			conn.retr(args)
		case "typeof":
			conn.setDataType(args)
		default:
			conn.respond(status502)
		}
	}
	if inputClient.Err() != nil {
		log.Print(inputClient.Err())
	}
}
