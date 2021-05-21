package formac

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestGetStruct(t *testing.T) {

	got, _ := getStruct(testMac)
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

func TestGetFormatted(t *testing.T) {
	cases := []struct {
		Format         string
		Want           string
		AssertFunction func(t testing.TB, got, want string)
	}{
		{
			Format: "plain",
			Want: `Cisco: 0000.5e00.5300
UnixExpanded: 00:00:5e:00:53:00
UnixCompact: 0:0:5e:0:53:0
EUI: 00-00-5E-00-53-00
Bare: 00005E005300
PgSQL: 00005e:005300`,
			AssertFunction: assertMACEqualString,
		},
		{
			Format:         "json",
			Want:           `{"Cisco": "0000.5e00.5300","UnixExpanded": "00:00:5e:00:53:00","UnixCompact": "0:0:5e:0:53:0","EUI": "00-00-5E-00-53-00","Bare": "00005E005300","PgSQL": "00005e:005300"}`,
			AssertFunction: assertMACEqualJSON,
		},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("Format to %s", test.Format), func(t *testing.T) {
			got, _ := GetFormatted(testMac, test.Format)
			test.AssertFunction(t, got, test.Want)
		})
	}
}

func assertMACEqualJSON(t testing.TB, got, want string) {
	t.Helper()

	var gotstruct, wantstruct macStruct

	json.Unmarshal([]byte(got), &gotstruct)
	json.Unmarshal([]byte(want), &wantstruct)

	if !reflect.DeepEqual(gotstruct, wantstruct) {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func assertMACEqualString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
