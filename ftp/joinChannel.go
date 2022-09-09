package ftp

import (
	"fmt"
)

/**
When you submit a command such as join
*/
func (c *Conn) joinChannel(args []string) {
	if len(args) != 1 {
		c.respond(status504)
		return
	}

	filtered := Filter(ChannelsAvailable, func(ch string) bool {
		return ch == args[0]
	})

	if len(filtered) != 1 {
		c.respond(fmt.Sprintf(status503, args[0]))
		return
	}

	fmt.Println(c.conn.RemoteAddr().String())
	c.dataUser = SetUser(c.conn.RemoteAddr().String(), args[0])

	//c.dataUser = dataUser
	/*workDir := filepath.Join(c.workDir, args[0])
	absPath := filepath.Join(c.rootDir, workDir)
	_, err := os.Stat(absPath)
	if err != nil {
		log.Print(err)
		c.respond(status550)
		return
	}
	c.workDir = workDir*/
	c.respond(status200)
}
