package formac

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
)

// All of the GetX functions *could* be implemented as
// methods on the internal macStruct type, but we want that
// type to remain internal but have a public way to return the
// representations.

type macStruct struct {
	Cisco        string
	UnixExpanded string
	UnixCompact  string
	EUI          string
	Bare         string
	PgSQL        string
}

// getStruct takes a mac address string and returns a macs struct
// with the supplied MAC address in various formats.
// This function is internal as we want a struct to DRY up
// rendering the different formats for representation in other
// convenience functions like GetPlain and GetJSON.
func getStruct(mac string) (macStruct, error) {
	var m macStruct
	hwaddr, err := net.ParseMAC(mac)
	if err != nil {
		return m, err
	}
	m = macStruct{
		Cisco:        FormatCisco(hwaddr),
		UnixExpanded: FormatUnixExpanded(hwaddr),
		UnixCompact:  FormatUnixCompact(hwaddr),
		EUI:          FormatEUI(hwaddr),
		Bare:         FormatBare(hwaddr),
		PgSQL:        FormatPgSQL(hwaddr),
	}
	return m, nil
}

// GetPlain takes a MAC address string and returns a multi-line,
// plaintext string of that MAC address in common formats.
func GetPlain(mac string) (string, error) {
	macs, err := getStruct(mac)
	if err != nil {
		return "", err
	}
	var sb strings.Builder
	fmt.Fprintf(&sb, "Cisco: %s\n", macs.Cisco)
	fmt.Fprintf(&sb, "UnixExpanded: %s\n", macs.UnixExpanded)
	fmt.Fprintf(&sb, "UnixCompact: %s\n", macs.UnixCompact)
	fmt.Fprintf(&sb, "EUI: %s\n", macs.EUI)
	fmt.Fprintf(&sb, "Bare: %s\n", macs.Bare)
	fmt.Fprintf(&sb, "PgSQL: %s", macs.PgSQL)
	return sb.String(), nil
}

// GetJSON takes a MAC address string and returns a non-prettified
// (single-line) JSON string of that MAC address in common formats.
// This *could* be implemented as JSON() method on
// the internal macStruct type, but we want that type to
// remain internal but have a public way to return the
// plaintext representation.
func GetJSON(mac string) (string, error) {
	macs, err := getStruct(mac)
	if err != nil {
		return "", err
	}
	b, err := json.Marshal(macs)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// GetFormatted takes a MAC address and output format as strings
// and returns the MAC address in various vendor formats in the
// requested output formats.
// Currently supports 'plain' and 'json' output formats. Outputs
// 'plain' by default.
func GetFormatted(mac, format string) (string, error) {
	var (
		s   string
		err error
	)

	switch format {
	case "json":
		s, err = GetJSON(mac)
	default:
		s, err = GetPlain(mac)
	}
	if err != nil {
		return "", err
	}
	return s, nil
}
