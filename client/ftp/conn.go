package ftp

import (
	log "github.com/sirupsen/logrus"
	"net"
)

type Conn struct {
	conn net.Conn
}

func NewConn(conn net.Conn) *Conn {
	logger := log.WithFields(log.Fields{
		"function": "NewConn",
	})
	logger.Info(" init config connection Client")
	return &Conn{
		conn: conn,
	}
}
