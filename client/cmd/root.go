package cmd

import (
	"bufio"
	"fmt"
	"github.com/lorenaggs/golang/client/ftp"
	log "github.com/sirupsen/logrus"
	"net"
	"os"
)

func Execute() {
	logger := log.WithFields(log.Fields{
		"function": "main",
	})
	logger.Info("Client is Ready")
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ftp.NewConn(conn)
	responseServer := make(chan string)
	for {
		go getResponseServer(conn, responseServer)
		go sendDataServer(conn)
		response := <-responseServer
		logger.Info(response)
	}
}

func getResponseServer(conn net.Conn, chIn chan<- string) {
	message, err := bufio.NewReaderSize(conn, 20999999).ReadString('|')
	if err != nil {
		log.Fatal(err)
	}
	chIn <- message
}

func sendDataServer(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	_, err = fmt.Fprint(conn, input)
	if err != nil {
		log.Error(err)
	}
}
