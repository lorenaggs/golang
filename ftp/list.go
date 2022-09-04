package ftp

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

/**
If the user didn’t provide a path argument, we list the contents of the current workDir.
*/
func (c *Conn) list(args []string) {
	var target string
	if len(args) > 0 {
		target = filepath.Join(c.rootDir, c.workDir, args[0])
	} else {
		target = filepath.Join(c.rootDir, c.workDir)
	}

	files, err := ioutil.ReadDir(target)  // returns each file in a directory
	if err != nil {
		log.Print(err)
		c.respond(status550)
		return
	}
	c.respond(status150)

	dataConn, err := c.dataConnect() // establish a second, temporary connection to the client,
	if err != nil {
		log.Print(err)
		c.respond(status425)
		return
	}
	defer dataConn.Close()

	for _, file := range files {
		_, err := fmt.Fprint(dataConn, file.Name(), c.EOL())
		if err != nil {
			log.Print(err)
			c.respond(status426)
		}
	}
	_, err = fmt.Fprintf(dataConn, c.EOL())
	if err != nil {
		log.Print(err)
		c.respond(status426)
	}

	c.respond(status226)
}
This is a more complex