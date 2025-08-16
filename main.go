package main

import (
	"log/slog"
	"os"

	"github.com/yashraj-n/tunnix/config"
	"github.com/yashraj-n/tunnix/network"
	"github.com/yashraj-n/tunnix/ssh"
)

func main() {
	config.SetupLogger()

	cliConfig := config.GetCliConfig()
	slog.Info("Attempting to connect to remote server", "server", cliConfig.RemoteIp)

	serverConn, err := ssh.AttemptConnection(cliConfig)

	if err != nil {
		slog.Error("Failed to connect to remote server", "error", err)
		os.Exit(1)
	}

	listener, err := ssh.CreateTunnel(serverConn, cliConfig.LocalPort, cliConfig.RemotePort)
	if err != nil {
		slog.Error("Failed to create tunnel", "error", err)
		os.Exit(1)
	}

	defer serverConn.Close()
	defer listener.Close()

	slog.Info("Connected to remote server")
	network.AcceptConnections(listener, cliConfig.LocalPort)
}
