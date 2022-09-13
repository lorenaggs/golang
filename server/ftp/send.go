package ftp

import log "github.com/sirupsen/logrus"

type file struct {
	channel    string
	fileName   string
	fileBase64 string
}

var filesShared []*file

func (c *Conn) send(args []string) {

	if len(args) != 3 {
		c.respond(status501)
		return
	}

	var channelSend = args[0]

	isValid := Filter(ChannelsAvailable, func(ch string) bool {
		return ch == channelSend
	})

	if len(isValid) != 1 {
		channelSend = c.dataUser.channel
	}

	file := &file{
		channel:    channelSend,
		fileName:   args[1],
		fileBase64: args[2],
	}

	filesShared = append(filesShared, file)

	log.Info(len(filesShared))
}
