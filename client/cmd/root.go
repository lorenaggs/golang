package cmd

import (
	"bufio"
	"fmt"
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
		fmt.Errorf("ERROR")
		log.Fatal(err)
		return
	}
	defer conn.Close()
	//ftp.NewConn(conn)
	responseServer := make(chan string)

	fmt.Printf("Commands accepted: \n channel \n join  \n\n\n")
	for {
		msg := getResponseServer(conn, responseServer)
		go sendDataServer(conn)
		fmt.Println(msg)
		/*response := <-responseServer
		logger.Info(response)*/
	}

}

func getResponseServer(conn net.Conn, chIn chan<- string) string {
	msg, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		return err.Error()
	}
	return string(msg)
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
