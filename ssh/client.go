package ssh

import (
	"fmt"

	"github.com/yashraj-n/tunnix/config"
	"golang.org/x/crypto/ssh"
)

func AttemptConnection(config config.CliConfig) (*ssh.Client, error) {
	sshConfig := &ssh.ClientConfig{
		User: config.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(config.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //TODO: Make this secure
	}

	serverConn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", config.RemoteIp, config.SSHPort), sshConfig)

	if err != nil {
		return nil, err
	}

	return serverConn, nil

}
