package config

import (
	"flag"
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

type CliConfig struct {
	Username   string
	Password   string
	RemoteIp   string
	SSHPort    int
	LocalPort  int
	RemotePort int
}

func SetupLogger() {
	slog.SetDefault(slog.New(tint.NewHandler(os.Stdout, &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: "15:04:05",
		NoColor:    false,
		AddSource:  true,
	})))
}

func GetCliConfig() CliConfig {
	var (
		password  = flag.String("password", "", "Password to use for SSH authentication")
		remoteIp  = flag.String("host", "", "Address to remote Server")
		localPort = flag.Int("port", 8080, "Local port to forward")
	)

	flag.Parse()

	if *password == "" || *remoteIp == "" || *localPort == 0 {
		slog.Error("Missing required flags", "password", *password, "remoteIp", *remoteIp, "localPort", *localPort)
		flag.PrintDefaults()
		os.Exit(1)
	}

	return CliConfig{
		Username:   "tunnix",
		Password:   *password,
		RemoteIp:   *remoteIp,
		SSHPort:    12000,
		LocalPort:  *localPort,
		RemotePort: 12001,
	}
}
