package ftp

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"path/filepath"
	"strings"
)

/*
*
If the user didn’t provide a path argument, we list the contents of the current workDir.
*/
func (c *Conn) list(args []string) {
	var target string
	responseFiles := []string{lbl_resp_list}
	if len(args) > 0 {
		target = filepath.Join(c.rootDir, c.workDir, args[0])
	} else {
		target = filepath.Join(c.rootDir, c.workDir)
	}

	files, err := ioutil.ReadDir(target) // returns each file in a directory
	if err != nil {
		log.Error(err)
		c.respond(status550)
		return
	}
	for _, file := range files {
		responseFiles = append(responseFiles, " -> "+file.Name())
	}
	c.respond(strings.Join(responseFiles, " : "))

}
