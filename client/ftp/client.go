package ftp

//
//import (
//	"bytes"
//	"fmt"
//	"io"
//	"net"
//)
//
//type Conn struct {
//	conn net.Conn
//}
//
//func (c *Conn) NewConn(conn net.Conn) error {
//	for {
//		msg := make([]byte, 2)
//		_, err := c.conn.Read(msg)
//
//		if err == io.EOF {
//			return nil
//		}
//
//		if err != nil {
//			return err
//		}
//
//		c.Handle(msg)
//	}
//}
//
//func (c *Conn) Handle(message []byte) {
//	cmd := bytes.ToUpper(bytes.TrimSpace(message))
//
//	switch string(cmd) {
//	case "RC":
//		c.receiveFile()
//	case "OK":
//		fmt.Println("OK")
//	default:
//		msg := make([]byte, 20)
//		c.conn.Read(msg)
//		fmt.Println(string(cmd) + string(msg))
//	}
//}
