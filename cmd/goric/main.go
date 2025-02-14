/*
Goric is a simple network toolkit command similar to nc.
It is a clone of "ric" https://github.com/cpmachado/ric
*/
package main

import (
	"flag"
	"fmt"
	"os"

	"go.cpmachado.pt/goric"
)

var Version = "0.5.0"

func main() {
	var host string = "localhost"
	var port int = 1337
	var v, n, w, u bool

	flag.Usage = displayUsage
	flag.BoolVar(&v, "v", false, "Display version and exit")
	flag.BoolVar(&n, "n", false, "Task 1: hostname")
	flag.BoolVar(&w, "w", false, "Task 2: nslook")
	flag.BoolVar(&u, "u", false, "Task 3,4,5: udp client to contact UDP echo server")
	flag.IntVar(&port, "p", port, "PORT")
	flag.StringVar(&host, "d", host, "HOSTNAME")
	flag.Parse()

	switch {
	case v:
		displayVersion()
	case n:
		goric.Hostname()
	case w:
		goric.Nslook(host)
	case u:
		goric.UDPClient(host, port)
	default:
		flag.Usage()
	}
}

func displayVersion() {
	fmt.Printf("goric-%s Copyright (C) 2025 cpmachado", Version)
}

func displayUsage() {
	prog := os.Args[0]
	fmt.Fprintf(flag.CommandLine.Output(),
		"%s is a WIP clone of 'ric'\n", prog)
	fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [OPTIONS]\n", prog)
	flag.PrintDefaults()
}
