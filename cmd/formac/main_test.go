package main

import (
	"bytes"
	"testing"
)

func TestFormatMac(t *testing.T) {
	macaddress := "00-00-5E-00-53-00"
	buffer := &bytes.Buffer{}
	FormatMac(buffer, macaddress)

	got := buffer.String()
	want := `Cisco: 0000.5e00.5300
UnixExpanded: 00:00:5e:00:53:00
UnixCompact: 0:0:5e:0:53:0
EUI: 00-00-5E-00-53-00
Bare: 00005E005300
PgSQL: 00005e:005300
`
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
