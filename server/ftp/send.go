package ftp

import (
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
)

type file struct {
	channel    string
	fileName   string
	fileBase64 string
}

var filesShared []*file

func (c *Conn) send(args []string) {
	if len(args) != 1 {
		c.respond(status501)
		return
	}

	path := filepath.Join(c.rootDir, c.workDir, args[0])
	file, err := os.Open(path)
	if err != nil {
		log.Error(err)
		c.respond(status550)
	}
	c.respond(status150)

	dataConn, err := c.dataConnect()
	if err != nil {
		log.Error(err)
		c.respond(status425)
	}
	defer dataConn.Close()

	_, err = io.Copy(dataConn, file)
	if err != nil {
		log.Error(err)
		c.respond(status426)
		return
	}
	io.WriteString(dataConn, c.EOL())
	c.respond(status226)
}
