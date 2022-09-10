package ftp

import (
	"fmt"
	"os"
	"path/filepath"
)

/**
When you submit a command such as join
*/
func (c *Conn) joinChannel(args []string) {
	if len(args) != 1 {
		c.respond(status504)
		return
	}
	channel := args[0]

	filtered := Filter(ChannelsAvailable, func(ch string) bool {
		return ch == channel
	})

	if len(filtered) != 1 {
		c.respond(fmt.Sprintf(status503, channel))
		return
	}

	fmt.Println(c.conn.RemoteAddr().String())
	c.dataUser = SetUser(c.conn.RemoteAddr().String(), channel)
	c.createFolder(channel)
	c.respond(status200)
}

func (c *Conn) hasUserChannel() bool {
	if c.dataUser == nil {
		c.respond(fmt.Sprintf(status503, ""))
		c.respond(lbl_question_channles)
	}
	return c.dataUser != nil
}

func (c *Conn) createFolder(channel string) {
	path := filepath.Join(c.rootDir, c.workDir, channel)
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		c.respond(err.Error())
		return
	}
}
