// Package ftp provides structs and functions for running an FTP server.
package ftp

//is our application’s router. It’s the workhorse that checks for incoming commands
//and matches them against the FTP routines that the server implements.
import (
	"bufio"
	"log"
	"strings"
)

// Serve scans incoming requests for valid commands and routes them to handler functions.
func Serve(c *Conn) {
	c.respond(status220) //The first thing we do upon entering Serve is to issue a 220 response to the client,

	s := bufio.NewScanner(c.conn) //To listen for incoming commands,
	for s.Scan() {
		input := strings.Fields(s.Text())
		if len(input) == 0 {
			continue
		}

		command, args := input[0], input[1:]  // you can see exactly what the client is sending
		log.Printf("<< %s %v", command, args) // you can see exactly what the client is sending

		switch command {
		case "cd":
			c.cwd(args)
		case "ls":
			c.list(args)
		case "port":
			c.port(args)
		case "user":
			c.user(args)
		case "exit":
			c.respond(status221)
			return
		case "RETR": // get //the client secretly sends a port
			c.retr(args)
		case "typeof":
			c.setDataType(args)
		default:
			c.respond(status502)
		}
	}
	if s.Err() != nil {
		log.Print(s.Err())
	}
}
