package main

import (
	"log/slog"

	"github.com/yashraj-n/tunnix/config"
	"github.com/yashraj-n/tunnix/ssh"
)

func main() {
	config.SetupLogger()

	cliConfig := config.GetCliConfig()

	slog.Info("Attempting to connect", "username", cliConfig.Username, "password", cliConfig.Password, "remoteIp", cliConfig.RemoteIp, "sshPort", cliConfig.SSHPort, "localPort", cliConfig.LocalPort, "remotePort", cliConfig.RemotePort)

	ssh.AttemptConnection(cliConfig)

	slog.Info("Connected to remote server")
	for {
	}

}
