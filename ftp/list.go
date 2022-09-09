package ftp

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

/*
*
If the user didn’t provide a path argument, we list the contents of the current workDir.
*/
func (c *Conn) list(args []string) {
	var target string
	responseFiles := []string{"test"}
	if len(args) > 0 {
		target = filepath.Join(c.rootDir, c.workDir, args[0])
	} else {
		target = filepath.Join(c.rootDir, c.workDir)
	}

	files, err := ioutil.ReadDir(target) // returns each file in a directory
	if err != nil {
		log.Print(err)
		c.respond(status550)
		return
	}
	for _, file := range files {
		responseFiles = append(responseFiles, file.Name())
	}
	c.respond(status150)
	c.respond(strings.Join(responseFiles, "\n"))
	fmt.Println(files)

}
