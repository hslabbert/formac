package formac

import (
	"net"
	"testing"
)

func TestFormatters(t *testing.T) {
	macaddress := "00-00-5E-00-53-00"

	var cases = []struct {
		Format    string
		Formatter func(net.HardwareAddr) string
		Want      string
	}{
		{Format: "Cisco", Formatter: FormatCisco, Want: "0000.5e00.5300"},
		{Format: "UnixExpanded", Formatter: FormatUnixExpanded, Want: "00:00:5e:00:53:00"},
		{Format: "UnixCompact", Formatter: FormatUnixCompact, Want: "0:0:5e:0:53:0"},
		{Format: "EUI", Formatter: FormatEUI, Want: "00-00-5E-00-53-00"},
		{Format: "Bare", Formatter: FormatBare, Want: "00005E005300"},
		{Format: "PgSQL", Formatter: FormatPgSQL, Want: "00005e:005300"},
	}

	hwmac, _ := net.ParseMAC(macaddress)
	for _, test := range cases {
		t.Run(test.Format, func(t *testing.T) {
			got := test.Formatter(hwmac)
			if got != test.Want {
				t.Errorf("got %v, wanted %v", got, test.Want)
			}
		})
	}
}

func TestGetStruct(t *testing.T) {
	macaddress := "00-00-5E-00-53-00"
	hwmac, _ := net.ParseMAC(macaddress)

	got := getStruct(hwmac)
	want := macs{
		Cisco:        "0000.5e00.5300",
		UnixExpanded: "00:00:5e:00:53:00",
		UnixCompact:  "0:0:5e:0:53:0",
		EUI:          "00-00-5E-00-53-00",
		Bare:         "00005E005300",
		PgSQL:        "00005e:005300",
	}

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
