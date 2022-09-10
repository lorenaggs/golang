package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
)

func main() {
	address := net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"), // Convierta la dirección IP de la cadena a tipo net.IP
		Port: 8080,
	}
	listener, err := net.ListenTCP("tcp", &address) // Crea un oyente del lado del servidor TCP
	if err != nil {
		log.Fatal(err) // Println + os.Exit(1)
	}
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatal(err) // Salir directamente después del error
		}
		fmt.Println("remote address:", conn.RemoteAddr())
		//go echo(conn)
	}
}
