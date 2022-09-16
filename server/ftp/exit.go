package ftp

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
)

func (c *Conn) exit() {

	id := c.conn.RemoteAddr().String()
	idNumber := strings.Split(id, "]:")

	msg := fmt.Sprintf(status224, idNumber[1])

	log.Warn(msg)

	c.respond(msg)
	c.conn.Close()

}
