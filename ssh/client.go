package ssh

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/yashraj-n/tunnix/config"
	"golang.org/x/crypto/ssh"
)

func AttemptConnection(config config.CliConfig) *ssh.Client {
	sshConfig := &ssh.ClientConfig{
		User: config.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(config.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //TODO: Make this secure
	}

	serverConn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", config.RemoteIp, config.SSHPort), sshConfig)

	if err != nil {
		slog.Error("Failed to connect to remote server", "error", err)
		os.Exit(1)
	}

	defer serverConn.Close()

	return serverConn

}
