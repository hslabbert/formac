package main

import (
	"fmt"
	"io"
	"net"
	"os"

	"github.com/hslabbert/formac/pkg/formac"
)

// FormatMac takes an io.Writer destination and a MAC address
// and writes the MAC address into the io.Writer in various
// common formats.
func FormatMac(out io.Writer, mac string) error {
	hwaddr, err := net.ParseMAC(mac)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "%s", formac.GetPlain(hwaddr))
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No arguments supplied; please provide a MAC address.")
		os.Exit(1)
	}
	macaddr := os.Args[1]
	err := FormatMac(os.Stdout, macaddr)
	if err != nil {
		fmt.Printf("Error encountered:\n%s\n", err)
	}
}
