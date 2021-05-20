package formac

import (
	"encoding/json"
	"net"
	"reflect"
	"testing"
)

func TestGetStruct(t *testing.T) {
	hwmac, _ := net.ParseMAC(testMac)

	got := getStruct(hwmac)
	want := macStruct{
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

func TestGetPlain(t *testing.T) {
	hwmac, _ := net.ParseMAC(testMac)

	got := GetPlain(hwmac)
	want := `Cisco: 0000.5e00.5300
UnixExpanded: 00:00:5e:00:53:00
UnixCompact: 0:0:5e:0:53:0
EUI: 00-00-5E-00-53-00
Bare: 00005E005300
PgSQL: 00005e:005300
`

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestGetJSON(t *testing.T) {
	hwmac, _ := net.ParseMAC(testMac)

	got, _ := GetJSON(hwmac)
	want := []byte(`{"Cisco":"0000.5e00.5300","UnixExpanded": "00:00:5e:00:53:00","UnixCompact":"0:0:5e:0:53:0","EUI": "00-00-5E-00-53-00","Bare":"00005E005300","PgSQL":"00005e:005300"}`)

	var gotstruct, wantstruct macStruct
	json.Unmarshal(got, &gotstruct)
	json.Unmarshal(want, &wantstruct)

	if !reflect.DeepEqual(gotstruct, wantstruct) {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
