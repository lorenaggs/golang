package ftp

import (
	"log"
	"os"
	"path/filepath"
)

/**
When you submit a command such as cd ../parent_folder to your FTP client, it sends that message to the server
*/
func (c *Conn) open(args []string) {
	if len(args) != 1 {
		c.respond(status501)
		return
	}

	workDir := filepath.Join(c.workDir, args[0])
	absPath := filepath.Join(c.rootDir, workDir)
	_, err := os.Stat(absPath)
	if err != nil {
		log.Print(err)
		c.respond(status550)
		return
	}
	c.workDir = workDir
	c.respond(status200)
}