package ftp

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
)

func (c *Conn) user(args []string) {
	c.respond(fmt.Sprintf(status230, strings.Join(args, " ")))

	log.Info("Hola")
}
