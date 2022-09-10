package ftp

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func (c *Conn) respond(s string) {
	_, err := fmt.Fprint(c.conn, s)
	if err != nil {
		log.Error(err)
	}
}
