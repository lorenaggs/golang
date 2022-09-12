package ftp

import (
	"fmt"
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
	addUser(c, channel)
}

func addUser(c *Conn, channel string) {
	var responseMessage = status201

	if c.dataUser != nil {
		responseMessage = status204
	}

	c.dataUser = SetUser(c.conn, c.conn.RemoteAddr().String(), channel)

	UsersConnected = append(UsersConnected, c.dataUser)

	c.respond(fmt.Sprintf(responseMessage, channel))
}

func (c *Conn) hasUserChannel() bool {
	if c.dataUser == nil {
		c.respond(lbl_question_channles)
	}
	return c.dataUser != nil
}
