package ftp

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

func Router(c *Client) {
	go func() {
		c.read()
	}()
	log.Warn("Please, first join into a channel, eg : join ch1")
	const menu = "\n\n M E N U: (select an option eg: join ch1) \n\n " +
		"join [ARG]\t\t\t\tSubscribe into another channel.  \n " +
		"send [PATH]\t\t\tSend file into channel.\n " +
		"chan\t\t\t\t\tList all channels \n " +
		"exit\t\t\t\t\tClose connection\n\n\n"
	fmt.Println(menu)
	reader := bufio.NewReader(os.Stdin)

	for {
		input, err := reader.ReadString('\n')

		if len(input) == 0 {
			continue
		}

		log.Debug(strings.Contains(input, join))
		if strings.Contains(input, send) {
			c.SendFile(input)
			continue
		}

		_, err = fmt.Fprint(c.conn, input)
		if err != nil {
			log.Error(err)
		}
	}

}
