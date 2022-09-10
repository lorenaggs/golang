package ftp

type dataType int

//keyword to set up the dataType constants ascii = 0 , binary = 1.
const (
	ascii dataType = iota
	binary
)

func (c *Conn) setDataType(args []string) {
	if len(args) == 0 {
		c.respond(status501)
	}

	switch args[0] {
	case "A": //ASCII
		c.dataType = ascii
	case "I": //Image/binary
		c.dataType = binary
	default:
		c.respond(status504)
		return
	}
	c.respond(status200)
}
