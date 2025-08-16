package network

import (
	"fmt"
	"io"
	"log/slog"
	"net"
)

// Accepts connections from remote server and handles them
func AcceptConnections(listener net.Listener, localPort int) {
	for {
		remoteConn, err := listener.Accept()
		if err != nil {
			slog.Error("Failed to accept connection", "error", err)
			continue
		}

		go handleConnection(remoteConn, localPort)
	}
}

// Handles a single connection from remote server
func handleConnection(remoteConn net.Conn, localPort int) {
	defer remoteConn.Close()
	localConn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", localPort))

	if err != nil {
		slog.Error("Failed to connect to local port", "error", err)
		return
	}
	defer localConn.Close()

	done := make(chan bool, 2) // bidirectional copy channel

	go func() {
		io.Copy(localConn, remoteConn)
		done <- true
	}()

	go func() {
		io.Copy(remoteConn, localConn)
		done <- true
	}()

	<-done

}
