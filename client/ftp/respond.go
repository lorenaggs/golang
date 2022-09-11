package ftp

const (
	join       = "join"
	channel    = "chan"
	listFolder = "list"
	send       = "send"
)
const (
	MaxBufferMb   = 10
	MaxBufferByte = MaxBufferMb * 1024 * 1024
)

/*func (c *Conn) respond(s string) {
	_, err := fmt.Fprint(c.conn, s)
	if err != nil {
		log.Error(err)
	}
}
*/
