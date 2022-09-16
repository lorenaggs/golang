// Package ftp provides structs and functions for running an FTP server.
package ftp

//is our application’s router. It’s the workhorse that checks for incoming commands
//and matches them against the FTP routines that the server implements.
import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"strings"
)

const (
	MaxBufferMb   = 10
	MaxBufferByte = MaxBufferMb * 1024 * 1024
)

// Serve scans incoming requests for valid commands and routes them to handler functions.
func Router(conn *Conn) {
	conn.respond(status220)                    //The first thing we do upon entering Serve is to issue a 220 response to the client,
	inputClient := bufio.NewScanner(conn.conn) //To listen for incoming commands,
	buffer := make([]byte, MaxBufferByte)
	inputClient.Buffer(buffer, MaxBufferByte)

	for inputClient.Scan() {
		input := strings.Fields(inputClient.Text())
		if len(input) == 0 {
			continue
		}

		command, args := input[0], input[1:] // you can see exactly what the client is sending
		log.WithFields(log.Fields{
			"args":    args[0:],
			"command": command,
		}).Info("The client is sending!")

		if command != "exit" && command != "join" && command != "chan" && !conn.hasUserChannel() {
			log.Warn("Client doesn't send command JOIN.")
			continue
		}

		switch command {
		case "join":
			conn.joinChannel(args)
		case "list":
			conn.list(args)
		case "port":
			conn.port(args)
		case "user": //ip that the client has
			conn.user(args)
		case "exit":
			conn.exit()
		case "send": // get //the client secretly sends a port
			conn.send(args)
		case "typeof":
			conn.setDataType(args)
		case "chan":
			conn.printChannels()
		default:
			conn.respond(status502)
		}
	}

	if inputClient.Err() != nil {
		log.Warn(inputClient.Err())
		log.Error(inputClient.Err())
	}

}
