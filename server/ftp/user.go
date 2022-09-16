package ftp

import (
	"fmt"
)

func (c *Conn) user(args []string) {
	//c.respond(fmt.Sprintf(status230, strings.Join(args, " ")))

	idClient := c.dataUser
	msg := fmt.Sprintf(status205, idClient.ip)
	c.respond(msg)

	//log.Info("Hola", c.dataUser)
}
