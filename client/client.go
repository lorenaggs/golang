package main

import (
	"flag"
	"fmt"
	"github.com/lorenaggs/golang/client/ftp"
	log "github.com/sirupsen/logrus"
	"net"
)

var host string
var port int
var rootDir string

func init() {
	flag.StringVar(&host, "host", "localhost", "port number")
	flag.IntVar(&port, "port", 8080, "port number")
	flag.StringVar(&rootDir, "rootDir", ".", "root directory")
	flag.Parse()
	log.SetFormatter(&log.TextFormatter{})
	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}
func main() {
	fmt.Printf(rootDir)
	logger := log.WithFields(log.Fields{
		"function": "main",
	})
	logger.Debug("Client is Ready")
	server := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.Dial("tcp", server)
	if err != nil {
		fmt.Errorf("ERROR")
		log.Fatal(err)
		return
	}
	defer conn.Close()
	handleConnections(conn, rootDir)

}

func handleConnections(c net.Conn, rootDir string) {
	ftp.Router(ftp.NewConn(c, rootDir))
}
