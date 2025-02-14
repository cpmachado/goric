// Package goric contains the implementations of each task
package goric

import (
	"fmt"
	"log/slog"
	"os"
)

// Hostname retrieves and prints host name
// Exits with non-zero error code should it have issues retrieving it
func Hostname() {
	hostname, err := os.Hostname()
	if err != nil {
		slog.Error("Hostname", slog.Any("error", err))
		os.Exit(1)
	}

	fmt.Printf("host name: %s\n", hostname)
}
