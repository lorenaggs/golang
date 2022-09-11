package main

import (
	"fmt"
	"github.com/lorenaggs/golang/client/ftp"
	log "github.com/sirupsen/logrus"
	"net"
)

func main() {
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
	handleConnections(conn)
	/*for {
		msg, err := ftp.GetResponseServer(conn)
		go ftp.SendDataServer(conn)
		fmt.Println(msg)
		if strings.TrimSpace(msg) == "# 200 Command okay." {
			break
		}
		if err != nil {
			break
		}
	}*/

}

func handleConnections(c net.Conn) {
	ftp.Router(ftp.NewConn(c))
}
