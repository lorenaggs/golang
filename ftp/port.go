package ftp

import log "github.com/sirupsen/logrus"

func (c *Conn) port(args []string) {
	if len(args) != 1 {
		c.respond(status501)
		return
	}
	dataPort, err := dataPortFromHostPort(args[0])
	if err != nil {
		log.Error(err)
		c.respond(status501)
		return
	}
	c.dataPort = dataPort
	c.respond(status200)
}
