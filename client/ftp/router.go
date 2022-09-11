package ftp

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func Router(c *Client) {
	logrus.Info("Please, first join into a channel, eg : join ch1")

	fmt.Printf("\n\n Commands accepted: \n * chan \n * join [ARG] \n\n\n")
	isChannelReady := c.readChannel()

	logrus.Info("You are in a channel!", isChannelReady)
	fmt.Printf("\n\n Commands accepted: \n * chan \n * join [ARG] \n * send [ARG] \n * list\n\n\n")

}
