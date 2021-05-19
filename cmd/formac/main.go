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
	fmt.Fprintf(out, "Cisco: %s\n", formac.FormatCisco(hwaddr))
	fmt.Fprintf(out, "UnixExpanded: %s\n", formac.FormatUnixExpanded(hwaddr))
	fmt.Fprintf(out, "UnixCompact: %s\n", formac.FormatUnixCompact(hwaddr))
	fmt.Fprintf(out, "EUI: %s\n", formac.FormatEUI(hwaddr))
	fmt.Fprintf(out, "Bare: %s\n", formac.FormatBare(hwaddr))
	fmt.Fprintf(out, "PgSQL: %s\n", formac.FormatPgSQL(hwaddr))
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
