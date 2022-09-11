package ftp

import (
	"fmt"
	"os"
	"path/filepath"
)

/*
*
When you submit a command such as join
*/
func (c *Conn) joinChannel(args []string) {
	if len(args) != 1 {
		c.respond(status504)
		return
	}
	channel := args[0]

	isValid := Filter(ChannelsAvailable, func(ch string) bool {
		return ch == channel
	})

	if len(isValid) != 1 {
		c.respond(fmt.Sprintf(status503, channel))
		return
	}
	createFolder(c, channel)

}
func createFolder(c *Conn, channel string) {
	var responseMessage = status200
	if c.dataUser != nil {
		responseMessage = status204
	}

	c.dataUser = SetUser(c.conn, c.conn.RemoteAddr().String(), channel)
	c.usersConnected = append(c.usersConnected, c.dataUser)

	path := filepath.Join(c.rootDir, c.workDir, channel)
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		c.respond(err.Error())
		return
	}
	c.respond(responseMessage)
}

func (c *Conn) hasUserChannel() bool {
	if c.dataUser == nil {
		c.respond(lbl_question_channles)
	}
	return c.dataUser != nil
}
