package main

import (
	"flag"
	"fmt"
	"github.com/GoesToEleven/GolangTraining/tcp/ftp"
	"log"
	"net"
	"path/filepath"
)

var port int
var rootDir string

/**
net.Listen  whit .Accept is similar to  http.ListenAndServe, here specify the protocol to use TCP and the address
*/
func main() {
	server := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", server)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
		}
		go handleConnection(conn)
	}
}

/**
handleConnection is our FTP connection, our concurrent FTP server,
every connection handle in its own gorutine, and clients not wait online to use the server

go handleConnection is our gorutine


*/
func handleConnection(c net.Conn) {
	defer c.Close()
	absPath, err := filepath.Abs(rootDir)
	if err != nil {
		log.Fatal(err)
	}
	ftp.Serve(ftp.NewConn(c, absPath))
}

func init() {
	flag.IntVar(&port, "port", 8080, "port number")
	flag.StringVar(&rootDir, "rootDir", "public", "root directory")
	flag.Parse()
}
