package ftp

import "net"

type dataUser struct {
	conn    net.Conn
	ip      string
	channel string
}

func SetUser(conn net.Conn, ip string, channel string) *dataUser {
	return &dataUser{
		conn:    conn,
		ip:      ip,
		channel: channel,
	}
}
