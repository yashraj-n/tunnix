package ssh

import (
	"fmt"
	"log/slog"
	"net"

	"golang.org/x/crypto/ssh"
)

func CreateTunnel(serverConn *ssh.Client, localPort int, remotePort int) (net.Listener, error) {
	listener, err := serverConn.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", remotePort))
	if err != nil {
		slog.Error("Failed to listen on remote server", "error", err)
		return nil, err
	}

	return listener, nil
}
