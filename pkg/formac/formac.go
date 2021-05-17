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