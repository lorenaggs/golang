package ftp

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
)

const (
	status150 = "150 File status okay; about to open data connection."
	status200 = "200 Command okay."
	status220 = "220 Service ready for new user. Please Join into a CHANNEL"
	status221 = "221 Service closing control connection."
	status226 = "226 Closing data connection. Requested file action successful."
	status230 = "230 User %s logged in, proceed."
	status425 = "425 Can't open data connection."
	status426 = "426 Connection closed; transfer aborted."
	status501 = "501 Syntax error in parameters or arguments."
	status502 = "502 Command not implemented."
	status503 = "503 Channel %s not implemented."
	status504 = "504 Command not implemented for that parameter."
	status550 = "550 Requested action not taken. File unavailable."
)
const (
	lbl_resp_list         = "Files founded :"
	lbl_question_channles = "List of available channels, please select one : eg: join ch1"
)

var ChannelsAvailable = []string{"ch1", "ch2", "ch3"}

/*
We EXTEND ftp.Conn with the method respond. respond takes a string, logs it,
and then copies it to its underlying net.Conn, which does the dirty work of sending the data to the client.
The only thing to watch out for here is the call to c.EOL, which addresses a quirk of the FTP standard.
*/
// respond copies a string to the client and terminates it with the appropriate FTP line terminator for the datatype.
func (c *Conn) respond(s string) {
	//log.Info(">> :: ", s)
	_, err := fmt.Fprint(c.conn, s, c.EOL())
	if err != nil {
		log.Error(err)
	}
}

func (c *Conn) printChannels() {
	resp := []string{lbl_question_channles}
	for _, channel := range ChannelsAvailable {
		resp = append(resp, " ■ "+channel)
	}
	c.respond(strings.Join(resp, "\n"))
}

// EOL returns the line terminator matching the FTP standard for the datatype.
func (c *Conn) EOL() string {
	switch c.dataType {
	case ascii:
		return "\r\n"
	case binary: //which is called “image”
		return "\n"
	default:
		return "\n"
	}
}