package goric

import (
	"bufio"
	"fmt"
	"log/slog"
	"net"
	"os"
)

// TCPClient is a simple TCP Client that sends "Hello!\n", to dest:port, and
// expects it to be an echo server that replies with something.
func TCPClient(host string, port int) {
	msg := []byte("Hello!\n")
	dest := fmt.Sprintf("%s:%d", host, port)

	conn, err := net.Dial("tcp", dest)
	if err != nil {
		slog.Error("TCP_CLIENT", slog.Any("error", err))
		os.Exit(1)
	}

	defer conn.Close()

	_, err = conn.Write(msg)
	if err != nil {
		slog.Error("TCP_CLIENT", slog.Any("error", err))
		os.Exit(1)
	}
	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		slog.Error("TCP_CLIENT", slog.Any("error", err))
		os.Exit(1)
	}

	fmt.Printf("echo: %q\n", string(data))
}
