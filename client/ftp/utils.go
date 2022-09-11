package ftp

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
	"os"
)

func GetResponseServer(conn net.Conn) (string, error) {
	msg, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		return "", err
	}
	if err == io.EOF {
		return "", nil
	}

	if err != nil {
		return "", err
	}
	return string(msg), nil
}

func SendDataServer(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	_, err = fmt.Fprint(conn, input)
	if err != nil {
		log.Error(err)
	}
}
