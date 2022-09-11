package ftp

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
	"os"
)

func GetResponseServer(conn net.Conn, test chan string) {
	msg, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		return
	}
	if err == io.EOF {
		return
	}

	if err != nil {
		return
	}
	test <- string(msg)
}

func SendDataServer(conn net.Conn) {
	log.Debug("Ingresa SendDataServer")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	_, err = fmt.Fprint(conn, input)
	if err != nil {
		log.Error(err)
	}
}
