package sshclient

import (
	"bytes"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh"
)

type Client struct {
	con *ssh.ClientConfig
}

func NewClient() (*Client, error) {

	if viper.GetString("ssh.privatekey") != "" {
		fmt.Println("")
	} else {
		return clientSSHpass()
	}
	return nil, nil

}

func clientSSHpass() (*Client, error) {
	config := &ssh.ClientConfig{
		User: viper.GetString("ssh.username"),
		Auth: []ssh.AuthMethod{
			ssh.Password(viper.GetString("ssh.password")),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return &Client{con: config}, nil
}

func (c *Client) ExecuteCmd(cmd, hostname string) string {

	var res string
	var stdoutBuf bytes.Buffer

	conn, err := ssh.Dial("tcp", hostname+":22", c.con)
	if err != nil {
		log.Printf("host is not available:%s\n", err.Error())
		return res
	}

	session, err := conn.NewSession()
	if err != nil {
		log.Println("can`t create session:", err.Error())
	}

	defer session.Close()

	session.Stdout = &stdoutBuf
	err = session.Run(cmd)
	if err != nil {
		log.Printf("can`t run cmd: %s", err.Error())
	}

	res = stdoutBuf.String()

	return res
}
