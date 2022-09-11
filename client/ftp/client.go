package ftp

import (
	log "github.com/sirupsen/logrus"
	"net"
	"strings"
)

type Client struct {
	conn net.Conn
}

func NewConn(conn net.Conn) *Client {
	return &Client{
		conn: conn,
	}
}

func (c *Client) readChannel() bool {
	for {
		msg, err := GetResponseServer(c.conn)
		go SendDataServer(c.conn)
		//this response is important to get after login user
		log.Info(msg)
		if strings.TrimSpace(msg) == "# 200 Command okay." {
			break
		}
		if err != nil {
			break
		}
	}
	return true
}
