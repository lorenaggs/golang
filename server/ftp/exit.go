package ftp

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
)

func (c *Conn) exit() {
	log.Info(len(UsersConnected))
	id := c.conn.RemoteAddr().String()
	idNumber := strings.Split(id, "]:")
	msg := fmt.Sprintf(status224, idNumber[1])

	newUserConn := Filter(UsersConnected, func(user *dataUser) bool {
		return user.ip != idNumber[1]
	})

	UsersConnected = newUserConn

	log.Info(len(UsersConnected))

	log.Warn(msg)
	c.respond(msg)
	c.conn.Close()

}
