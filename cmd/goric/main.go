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

func init() {
	var v bool

	flag.Usage = displayUsage
	flag.BoolVar(&v, "v", false, "Display version and exit")
	flag.Parse()

	if v {
		displayVersion()
		os.Exit(0)
	}
	if flag.NArg() > 1 {
		err := fmt.Errorf("goric only accepts 1 argument at most")
		slog.Error("Init", slog.Any("error", err))
		os.Exit(1)
	}
}

func main() {
	var dest string

	goric.Hostname()
	if flag.NArg() == 1 {
		dest = flag.Arg(0)
		goric.Nslook(dest)
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
