package ftp

//This is our purpose-built representation of a connection to our FTP server. Notice that it wraps the original net.Conn
import "net"

// Conn represents a connection to the FTP server
type Conn struct {
	respChannel chan string
	conn        net.Conn  //will do all the direct communication with the client for us.
	dataType    dataType  //zero value 0
	dataPort    *dataPort // zero value nil
	rootDir     string    //we specified for the server (the place where public files will live)
	workDir     string    //the current working directory for the connection
}

// NewConn returns a new FTP connection
func NewConn(conn net.Conn, rootDir string, respChannel chan string) *Conn {
	return &Conn{
		respChannel: respChannel,
		conn:        conn,
		rootDir:     rootDir,
		workDir:     "/",
	}
}
