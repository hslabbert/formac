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

// getStruct takes a net.HardwareAddr and returns a macs struct
// with the supplied net.HardwareAddr in various formats.
// This function is internal as we want a struct to DRY up
// rendering the different formats for representation in other
// convenience functions like GetPlain and GetJSON.
func getStruct(hwaddr net.HardwareAddr) macStruct {
	m := macStruct{
		Cisco:        FormatCisco(hwaddr),
		UnixExpanded: FormatUnixExpanded(hwaddr),
		UnixCompact:  FormatUnixCompact(hwaddr),
		EUI:          FormatEUI(hwaddr),
		Bare:         FormatBare(hwaddr),
		PgSQL:        FormatPgSQL(hwaddr),
	}
	return m
}

// GetPlain takes a net.HardwareAddr and returns a multi-line,
// plaintext string of that MAC address in common formats.
func GetPlain(hwaddr net.HardwareAddr) string {
	macs := getStruct(hwaddr)
	var sb strings.Builder
	fmt.Fprintf(&sb, "Cisco: %s\n", macs.Cisco)
	fmt.Fprintf(&sb, "UnixExpanded: %s\n", macs.UnixExpanded)
	fmt.Fprintf(&sb, "UnixCompact: %s\n", macs.UnixCompact)
	fmt.Fprintf(&sb, "EUI: %s\n", macs.EUI)
	fmt.Fprintf(&sb, "Bare: %s\n", macs.Bare)
	fmt.Fprintf(&sb, "PgSQL: %s\n", macs.PgSQL)
	return sb.String()
}

// GetJSON takes a net.HardwareAddr and returns a non-prettified
// (single-line) JSON string of that MAC address in common formats.
// This *could* be implemented as JSON() method on
// the internal macStruct type, but we want that type to
// remain internal but have a public way to return the
// plaintext representation.
func GetJSON(hwaddr net.HardwareAddr) ([]byte, error) {
	macs := getStruct(hwaddr)
	b, err := json.Marshal(macs)
	if err != nil {
		return nil, err
	}
	return b, nil
}
