package ftp

type dataUser struct {
	ip      string
	channel string
}

func SetUser(ip string, channel string) *dataUser {
	return &dataUser{
		ip:      ip,
		channel: channel,
	}
}
