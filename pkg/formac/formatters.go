package formac

import (
	"net"
	"strings"
)

// FormatCisco takes a net.HardwareAddress and returns a Cisco-formatted
// MAC address.
func FormatCisco(hwaddr net.HardwareAddr) string {
	mac := hwaddr.String()
	mac = strings.Replace(mac, ":", "", -1)

	var sb strings.Builder
	for i := 0; i < 12; {
		if i != 0 {
			sb.WriteString(".")
		}
		sb.WriteString(mac[i : i+4])
		i += 4
	}

	return sb.String()
}

// FormatUnixExpanded takes a net.HardwareAddress and returns a Unix-formatted
// MAC address.
func FormatUnixExpanded(hwaddr net.HardwareAddr) string {
	return hwaddr.String()
}

// FormatUnixCompact takes a net.HardwareAddress and returns a Unix-formatted
// MAC address with extraneous zeroes ("0") stripped.
func FormatUnixCompact(hwaddr net.HardwareAddr) string {
	mac := hwaddr.String()
	var sb strings.Builder
	for i := 0; i < len(mac); i += 3 {
		if mac[i] == 48 {
			sb.WriteString(mac[i+1 : i+2])
		} else {
			sb.WriteString(mac[i : i+2])
		}
		if i+2 < len(mac) {
			sb.WriteString(":")
		}
	}

	return sb.String()
}

// FormatPgSQL takes a net.HardwareAddress and returns a PostgresQL-formatted
// MAC address.
func FormatPgSQL(hwaddr net.HardwareAddr) string {
	mac := hwaddr.String()
	macStripped := strings.Replace(mac, ":", "", -1)

	var sb strings.Builder
	sb.WriteString(macStripped[0:6])
	sb.WriteString(":")
	sb.WriteString(macStripped[6:12])

	return sb.String()
}

// FormatBare takes a net.HardwareAddress and returns a "bare"
// MAC address, i.e. capital hex characters only with no delimiters.
func FormatBare(hwaddr net.HardwareAddr) string {
	mac := hwaddr.String()
	macStripped := strings.Replace(mac, ":", "", -1)

	return strings.ToUpper(macStripped)
}

// FormatEUI takes a net.HardwareAddress and returns an EUI-formatted
// MAC address, i.e. capital hex characters with hyphen ("-") separating each
// byte's worth of hex characters.
func FormatEUI(hwaddr net.HardwareAddr) string {
	mac := hwaddr.String()
	macEUI64Lower := strings.Replace(mac, ":", "-", -1)

	return strings.ToUpper(macEUI64Lower)
}
