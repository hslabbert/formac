package formac

import (
	"bytes"
	"fmt"
	"net"
	"reflect"
	"testing"
)

const testRegistryString = `Registry,Assignment,Organization Name,Organization Address
MA-L,000001,XEROX CORPORATION,M/S 105-50C WEBSTER NY US 14580
MA-L,00005E,"ICANN, IANA Department",INTERNET ASS'NED NOS.AUTHORITY Los Angeles CA US 90094-2536
MA-L,54A493,IEEE Registration Authority,445 Hoes Lane Piscataway NJ US 08554`

func TestGetManufacturer(t *testing.T) {
	cases := []struct {
		Registry string
		MAC      string
		Want     string
	}{
		{
			Registry: "oui",
			MAC:      testMac,
			Want:     "ICANN, IANA Department",
		},
		{
			Registry: "cid",
			MAC:      "BA-55-EC-00-53-01",
			Want:     "IEEE 802.15",
		},
		{
			Registry: "iab",
			MAC:      "00-50-C2-4A-40-01",
			Want:     "IEEE P1609 WG",
		},
		{
			Registry: "mam",
			MAC:      "74-1A-E0-90-53-01",
			Want:     "Private",
		},
		{
			Registry: "mas",
			MAC:      "70-B3-D5-01-B0-01",
			Want:     "AUDI AG",
		},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("Retrieve a MAC vendor entry from the %s registry", test.Registry), func(t *testing.T) {
			hwmac, _ := net.ParseMAC(test.MAC)
			got := GetManufacturer(hwmac)

			if got != test.Want {
				t.Errorf("got %s, wanted %s", got, test.Want)
			}
		})
	}
}

func TestLoadCSVRegistry(t *testing.T) {

	regReader := bytes.NewReader([]byte(testRegistryString))

	got := loadCSVRegistry(regReader)
	want := macManufRegistryMap{
		"000001": "XEROX CORPORATION",
		"00005E": "ICANN, IANA Department",
		"54A493": "IEEE Registration Authority",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestLoadCSVRegistries(t *testing.T) {
	cases := []struct {
		Name string
		File string
	}{
		{
			Name: "oui",
			File: "data/oui.csv",
		},
		{
			Name: "cid",
			File: "data/cid.csv",
		},
		{
			Name: "iab",
			File: "data/iab.csv",
		},
		{
			Name: "mam",
			File: "data/mam.csv",
		},
		{
			Name: "mas",
			File: "data/oui36.csv",
		},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("try loading %s CSV file", test.Name), func(t *testing.T) {
			_, err := csvfs.ReadFile(test.File)
			if err != nil {
				t.Errorf("Unable to read file %s", test.File)
			}
		})
	}
}
