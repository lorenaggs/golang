package ftp

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

func Router(c *Client) {
	inputClient := bufio.NewReader(os.Stdin)
	log.Info("Please, first join into a channel, eg : join ch1")

	fmt.Printf("\n\n Commands accepted: \n * %s # List available channels\n * %s [ARG] \n\n\n", channel, join)
	isChannelReady := c.readChannel()

	if !isChannelReady {
		return
	}

	log.Warn("You are in a channel!")
	fmt.Printf("\n\n M E N U: (select an option eg: 1) \n\n 1. Send file into channel. \n 2. Subscribe into another channel. \n 3. List all channels \n 4. Exit\n\n\n")

	for {
		log.Debug("Ingresa menu")
		option, _ := inputClient.ReadString('\n')
		option = strings.TrimSpace(option)
		//log.Debug("opcion ->", option)
		switch option {
		case "1":
			log.Info("Input the channel and the path file eg: ch1 /user/home/photo.png")
			c.SendFile()
			continue
		case "2":
		case "3":
		case "4":
		}
	}

}
