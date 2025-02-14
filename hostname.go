package main

import (
	"fmt"
	"log/slog"
	"os"
)

func Hostname() {
	hostname, err := os.Hostname()
	if err != nil {
		slog.Error("Hostname", slog.Any("error", err))
		os.Exit(1)
	}

	fmt.Printf("host name: %s\n", hostname)
}
