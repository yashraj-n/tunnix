package ssh

import (
	"errors"
	"fmt"

	"github.com/yashraj-n/tunnix/config"
	"golang.org/x/crypto/ssh"
)

// Attempts to connect to remote server and returns client, motd, and error
func AttemptConnection(config config.CliConfig) (*ssh.Client, string, error) {
	sshConfig := &ssh.ClientConfig{
		User: config.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(config.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //TODO: Make this secure
	}

	serverConn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", config.RemoteIp, config.SSHPort), sshConfig)

	if err != nil {
		return nil, "", errors.New("failed to connect to remote server: " + err.Error())
	}

	session, err := serverConn.NewSession()
	if err != nil {
		return nil, "", errors.New("failed to create session: " + err.Error())
	}

	motd, err := session.CombinedOutput("cat /etc/motd")
	if err != nil {
		return nil, "", errors.New("failed to get MOTD: " + err.Error())
	}

	session.Close()
	return serverConn, string(motd), nil

}
