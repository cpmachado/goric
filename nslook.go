package goric

import (
	"fmt"
	"log/slog"
	"net"
	"os"
)

// Nslook prints CNAME of a given destination and IP's associated with it
func Nslook(dest string) {
	cname, err := net.LookupCNAME(dest)
	if err != nil {
		slog.Error("nslook", slog.Any("error", err))
		os.Exit(1)
	}

	fmt.Printf("CNAME: %s\n", cname)

	ips, err := net.LookupIP(dest)
	if err != nil {
		slog.Error("nslook", slog.Any("error", err))
		os.Exit(1)
	}

	fmt.Println("IP:")
	for _, ip := range ips {
		fmt.Printf("- %s\n", ip.String())
	}
}
