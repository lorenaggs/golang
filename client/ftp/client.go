package ftp

import (
	"bufio"
	"encoding/base64"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
)

type Client struct {
	conn    net.Conn //will do all the direct communication with the server for us.
	rootDir string   //we specified for the server (the place where public files will live)
	workDir string   //the current working directory for the connection
}

func NewConn(conn net.Conn, rootDir string) *Client {
	return &Client{
		conn:    conn,
		rootDir: rootDir,
		workDir: "/",
	}
}

func (c *Client) readChannel() bool {
	for {
		test := make(chan string)
		go GetResponseServer(c.conn, test)
		go SendDataServer(c.conn)

		msg := <-test
		//this response is important to get after login user
		log.Info(msg)
		if strings.TrimSpace(msg) == "# 200 Command okay." {
			return true
		}
	}
	return false
}

func (c *Client) SendFile() {
	for {
		reader := bufio.NewReader(os.Stdin)
		// ReadString will block until the delimiter is entered
		input, _ := reader.ReadString('\n')
		//input, err := reader.ReadString('\n')
		channelPath := strings.Fields(input)
		if len(channelPath) < 2 {
			continue
		}

		//channel, filePath := channelPath[0], channelPath[1]
		_, filePath := channelPath[0], channelPath[1]
		_, error := os.Stat(filePath) //validate if file exist

		// check if error is "file not exists"
		if os.IsNotExist(error) {
			log.Errorf("%s File does not exist. send again eg: ch1 /user/home/photo.png \n", filePath)
			continue
		}

		//tranform byte to string
		base64File(filePath)

		//command := fmt.Sprintf("%s %s", send, channel)

		//fmt.Println(command)
		//_, err = c.conn.Write([]byte(command))
		//if err != nil {
		//	log.Error(err)
		//}
	}
}

func base64File(filePath string) (string, error) {
	fileOpen, err := os.Open(filePath)
	if err != nil {
		log.Errorf("Error open file  %s \n", filePath)
	}
	defer fileOpen.Close()
	reader := bufio.NewReaderSize(fileOpen, MaxBufferByte)
	fileByte, err := io.ReadAll(reader)
	fmt.Println(fileByte)
	if err != nil {
		log.Error("Error reading file", err.Error())
	}
	if len(fileByte) > MaxBufferByte {
		log.Errorf("%s File is higher that permited %d \n", filePath, MaxBufferMb)
	}
	fileBase64 := base64.StdEncoding.EncodeToString(fileByte)
	fmt.Println(fileBase64)

	return "", nil
}
func createFolder(c *Client, channel string) {
	path := filepath.Join(c.rootDir, c.workDir, c.conn.RemoteAddr().String(), channel)
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Info(":: folder created ")
		return
	}
}
