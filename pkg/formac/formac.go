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

// FormatPgSQL takes a net.HardwareAddress and returns a PostgresQL-formatted
// MAC address.
func FormatPgSQL(hwaddr net.HardwareAddr) string {
	mac := hwaddr.String()
	mac = strings.Replace(mac, ":", "", -1)
	var sb strings.Builder
	sb.WriteString(mac[0:6])
	sb.WriteString(":")
	sb.WriteString(mac[6:12])
	return sb.String()
}
