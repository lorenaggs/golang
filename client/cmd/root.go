package cmd

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
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
	for {
		//go getResponseServer(conn, responseServer)
		go sendDataServer(conn)
		response := <-responseServer
		logger.Info(response)
	}
}

func getResponseServer(conn net.Conn, chIn chan<- string) {
	buf := make([]byte, 0, 4096) // big buffer
	tmp := make([]byte, 2556)    // using small tmo buffer for demonstrating
	n, err := conn.Read(tmp)
	if err != nil {
		if err != io.EOF {
			fmt.Println("read error:", err)
		}
		return
	}
	//fmt.Println("got", n, "bytes.")
	buf = append(buf, tmp[:n]...)
	fmt.Println("total size:", len(buf))
	fmt.Println("SERVER -->:", string(buf))
	/*	message, err := bufio.NewReaderSize(conn, controllers.MAX_BUFFER).ReadString('#')
		if err != nil {
			log.Error(err)
		}
		chIn <- message*/
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
