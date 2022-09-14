package main

import (
	"flag"
	"fmt"
	"github.com/lorenaggs/golang/server/ftp"
	log "github.com/sirupsen/logrus"
	"net"
	"path/filepath"
)

var port int
var rootDir string

func init() {
	flag.IntVar(&port, "port", 8080, "port number")
	flag.StringVar(&rootDir, "rootDir", "public", "root directory")
	flag.Parse()
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	log.SetLevel(log.DebugLevel)
}

// net.Listen  whit .Accept is similar to  http.ListenAndServe, here specify the protocol to use TCP and the address
func main() {
	logger := log.WithFields(log.Fields{
		"function": "main",
	})
	logger.Info("Server is Ready")
	server := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", server)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Error(err)
		}
		go handleConnection(conn)
	}
}

// handleConnection is our FTP connection, our concurrent FTP server,
// every connection handle in its own gorutine, and clients not wait online to use the server
// go handleConnection is our gorutine
func handleConnection(c net.Conn) {
	defer c.Close()
	absPath, err := filepath.Abs(rootDir)
	if err != nil {
		log.Fatal(err)
	}
	ftp.Router(ftp.NewConn(c, absPath))
}
