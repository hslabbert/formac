package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hslabbert/formac/pkg/formac"
)

func flagUsage() {
	w := flag.CommandLine.Output()
	fmt.Fprintf(w, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()

	fmt.Fprintf(w, "  string\n")
	fmt.Fprintf(w, "        The MAC address to parse and format\n")
}

func main() {
	flag.Usage = flagUsage
	formatPtr := flag.String("format", "plain", "MAC output format")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("No arguments supplied; please provide a MAC address.")
		os.Exit(1)
	}

	macaddr := flag.Arg(0)

	macs, geterr := formac.GetFormatted(macaddr, *formatPtr)
	if geterr != nil {
		fmt.Printf("Error encountered:\n%s\n", geterr)
	}

	_, fmterr := fmt.Fprintf(os.Stdout, "%s\n", macs)
	if fmterr != nil {
		fmt.Printf("Error encountered:\n%s\n", fmterr)
	}
}
