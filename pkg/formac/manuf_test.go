package formac

import (
	"fmt"
	"net"
	"reflect"
	"testing"
	"testing/fstest"
)

const testRegistryString = `Registry,Assignment,Organization Name,Organization Address
MA-L,000001,XEROX CORPORATION,M/S 105-50C WEBSTER NY US 14580
MA-L,00005E,"ICANN, IANA Department",INTERNET ASS'NED NOS.AUTHORITY Los Angeles CA US 90094-2536
MA-L,54A493,IEEE Registration Authority,445 Hoes Lane Piscataway NJ US 08554`

func TestGetManufacturer(t *testing.T) {
	hwmac, _ := net.ParseMAC(testMac)
	got := GetManufacturer(hwmac)
	want := "ICANN, IANA Department"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestLoadCSVRegistry(t *testing.T) {

	m := fstest.MapFS{
		"testregistry.csv": {
			Data: []byte(testRegistryString),
		},
	}
	regdata, _ := m.ReadFile("testregistry.csv")

	got := loadCSVRegistry(&regdata)
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
