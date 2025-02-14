package goric

import (
	"bufio"
	"fmt"
	"log/slog"
	"net"
	"os"
)

// UDPClient is a simple UDP Client that sends "Hello!\n", to dest:port, and
// expects it to be an echo server that replies with something.
func UDPClient(host string, port int) {
	msg := []byte("Hello!\n")
	dest := fmt.Sprintf("%s:%d", host, port)

	udpAddr, err := net.ResolveUDPAddr("udp", dest)
	if err != nil {
		slog.Error("UDP_CLIENT", slog.Any("error", err))
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		slog.Error("UDP_CLIENT", slog.Any("error", err))
		os.Exit(1)
	}

	_, err = conn.Write(msg)
	if err != nil {
		slog.Error("UDP_CLIENT", slog.Any("error", err))
		os.Exit(1)
	}
	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		slog.Error("UDP_CLIENT", slog.Any("error", err))
		os.Exit(1)
	}

	fmt.Printf("echo: %q\n", string(data))
	dest = conn.LocalAddr().String()
	fmt.Printf("sent by [%s]\n", dest)
}
