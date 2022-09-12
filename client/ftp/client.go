package ftp

import (
	"bufio"
	"encoding/base64"
	"errors"
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

func (c *Client) read() {
	for {
		response := make(chan string)
		go GetResponseServer(c.conn, response)
		msg := <-response
		//this response is important to get after login user
		log.Info(msg)
	}
}

func (c *Client) SendFile(input string) {
	sendCommand := strings.Fields(input)
	if len(sendCommand) > 1 && len(sendCommand) <= 3 {
		log.Error("Command Invalid, eg: send [channel] [path file]")
		return
	}
	//send /folde/file.txt
	//send ch1 /folde/file.txt
	var channel, filePath string
	_, channelPath := sendCommand[0], sendCommand[1:]
	log.Debug(channelPath)
	filePath = channelPath[0]
	if len(channelPath) == 2 {
		channel, filePath = channelPath[0], channelPath[1]
	}
	if channel == "" {
		channel = "nochan"
	}

	_, error := os.Stat(filePath) //validate if file exist

	// check if error is "file not exists"
	if os.IsNotExist(error) {
		log.Errorf("%s File does not exist. send again eg: send ch1 /user/home/photo.png \n", filePath)
		return
	}

	//tranform byte to string
	fileBase64, fileName, err := base64File(filePath)
	if err != nil {
		return
	}

	command := fmt.Sprintf("%s %s %s %s \n", send, channel, fileName, fileBase64)
	log.Debug(command)
	_, err = c.conn.Write([]byte(command))
	if err != nil {
		panic(err)
	}

}

func base64File(filePath string) (string, string, error) {
	fileOpen, err := os.Open(filePath)
	if err != nil {
		log.Errorf("Error open file  %s \n", filePath)
	}
	defer fileOpen.Close()
	reader := bufio.NewReaderSize(fileOpen, MaxBufferByte)
	fileByte, err := io.ReadAll(reader)
	log.Debug(fileByte)
	if err != nil {
		log.Error("Error reading file", err.Error())
		return "", "", err
	}
	fileInfo, err := fileOpen.Stat()
	if err != nil {
		log.Error("Error reading file", err.Error())
		return "", "", err
	}
	if len(fileByte) > MaxBufferByte {
		log.Errorf("%s File is higher that permited %d \n", filePath, MaxBufferMb)
		err = errors.New("File is higher that permited ")
		return "", "", err
	}
	fileBase64 := base64.StdEncoding.EncodeToString(fileByte)
	logger := log.WithFields(log.Fields{
		"function":     "base64File",
		"len-fileByte": len(fileByte),
		"fileInfo":     fileInfo.Size(),
	})
	logger.Debug("file information")
	return fileBase64, fileInfo.Name(), nil
}
func createFolder(c *Client, channel string) {
	path := filepath.Join(c.rootDir, c.workDir, c.conn.RemoteAddr().String(), channel)
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Info(":: folder created ")
		return
	}
}
