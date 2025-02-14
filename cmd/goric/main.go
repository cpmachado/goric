/*
Goric is a simple network toolkit command similar to nc.
It is a clone of "ric" https://github.com/cpmachado/ric
*/
package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"go.cpmachado.pt/goric"
)

var Version = "0.1.1"

func main() {
	var host string = "localhost"
	var v, n, w bool

	flag.Usage = displayUsage
	flag.BoolVar(&v, "v", false, "Display version and exit")
	flag.BoolVar(&n, "n", false, "Task 1: hostname")
	flag.BoolVar(&w, "w", false, "Task 2: nslook")
	flag.Parse()

	switch flag.NArg() {
	case 0:
	case 1:
		host = flag.Arg(0)
	default:
		err := fmt.Errorf("goric only accepts 1 argument at most")
		slog.Error("Init", slog.Any("error", err))
		os.Exit(1)
	}

	switch {
	case v:
		displayVersion()
	case n:
		goric.Hostname()
	case w:
		goric.Nslook(host)
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
		"%s is a WIP clone of 'ric', currently prints hostname\n", prog)
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", prog)
	flag.PrintDefaults()
}
