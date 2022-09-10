package ftp

//This is our purpose-built representation of a connection to our FTP server. Notice that it wraps the original net.Conn
import (
	log "github.com/sirupsen/logrus"
	"net"
)

// Conn represents a connection to the FTP server
type Conn struct {
	conn           net.Conn    //will do all the direct communication with the client for us.
	dataType       dataType    //zero value 0
	dataPort       *dataPort   // zero value nil
	rootDir        string      //we specified for the server (the place where public files will live)
	workDir        string      //the current working directory for the connection
	dataUser       *dataUser   //handle information about user connection
	usersConnected []*dataUser //save all users connected
}

// NewConn returns a new FTP connection
func NewConn(conn net.Conn, rootDir string) *Conn {
	logger := log.WithFields(log.Fields{
		"function": "NewConn",
	})
	logger.Info("init config connection")
	return &Conn{
		conn:    conn,
		rootDir: rootDir,
		workDir: "/",
	}
}
