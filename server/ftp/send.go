package ftp

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

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

	//todo : envio rápido tomar objeto file,y buscar todos los usuarios que pertenecen al canal del archivo que se recibe en el servidor
	//todo : crear una go rutina,
	//todo : crear canal que avise cuando se ha enviado el archivo

	log.Info(len(filesShared))

	if filesShared != nil {
		c.respond(fmt.Sprintf(status222, channelSend))
	}

	userByChannel := Filter(UsersConnected, func(user *dataUser) bool {

		return user.ip != c.dataUser.ip && user.channel == channelSend
	})

	for _, user := range userByChannel {

		command := fmt.Sprintf(status223, user.channel, file.fileName, file.fileBase64, user.ip)
		RespondUsers(user.conn, command)
		log.Debug(command)
	}

	log.Debug("A estos usuarios voy a enviar los archivos", len(userByChannel))

}
