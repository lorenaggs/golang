package ftp

import log "github.com/sirupsen/logrus"

const (
	join = "join"
	send = "send"
)
const (
	MaxBufferMb   = 10
	MaxBufferByte = MaxBufferMb * 1024 * 1024
)

func (c *Client) request(command string) error {
	log.Debug("request: " + command)
	_, err := c.conn.Write([]byte(command))
	if err != nil {
		return err
	}
	return nil
}
