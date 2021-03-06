package formac

import (
	"bytes"
	"embed"
	"encoding/csv"
	"io"
	"net"
)

//go:embed data/*
var csvfs embed.FS

type macManufRegistry struct {
	PrefixLength int
	Registry     macManufRegistryMap
}

type macManufRegistryMap map[macPrefix]macManufacturer

type macPrefix string
type macManufacturer string

type registrySearchResult struct {
	Manufacturer macManufacturer
	Found        bool
}

var macManufRegistries = map[string]macManufRegistry{
	"oui": {
		PrefixLength: 6,
		Registry:     loadCSVRegistry(newFileReader(csvfs, "data/oui.csv")),
	},
	"cid": {
		PrefixLength: 6,
		Registry:     loadCSVRegistry(newFileReader(csvfs, "data/cid.csv")),
	},
	"iab": {
		PrefixLength: 9,
		Registry:     loadCSVRegistry(newFileReader(csvfs, "data/iab.csv")),
	},
	"mam": {
		PrefixLength: 7,
		Registry:     loadCSVRegistry(newFileReader(csvfs, "data/mam.csv")),
	},
	"mas": {
		PrefixLength: 9,
		Registry:     loadCSVRegistry(newFileReader(csvfs, "data/oui36.csv")),
	},
}

// GetManufacturer takes a net.HardwareAddr and returns the
// manufacturer to whom that MAC address is registered.
func GetManufacturer(mac net.HardwareAddr) string {

	bareMAC := macPrefix(FormatBare(mac))
	manuf := searchRegistry(bareMAC, macManufRegistries["oui"])
	if manuf.Found && manuf.Manufacturer != "IEEE Registration Authority" {
		return string(manuf.Manufacturer)
	}

	resultChannel := make(chan registrySearchResult)
	registries := []string{"cid", "iab", "mam", "mas"}

	for _, registry := range registries {
		go func(r string) {
			resultChannel <- searchRegistry(bareMAC, macManufRegistries[r])
		}(registry)
	}

	for i := 0; i < len(registries); i++ {
		r := <-resultChannel
		if r.Found {
			return string(r.Manufacturer)
		}
	}
	return "Not found"
}

func loadCSVRegistry(regReader io.Reader) macManufRegistryMap {
	csvReader := csv.NewReader(regReader)
	rawCSVdata, _ := csvReader.ReadAll()

	r := make(macManufRegistryMap)
	for lineNum, record := range rawCSVdata {
		if lineNum == 0 {
			continue
		}
		r[macPrefix(record[1])] = macManufacturer(record[2])
	}
	return r
}

func newFileReader(fs embed.FS, name string) io.Reader {
	b, _ := fs.ReadFile(name)
	return bytes.NewReader(b)
}

func searchRegistry(mac macPrefix, r macManufRegistry) registrySearchResult {
	pfx := mac[:r.PrefixLength]
	manuf, ok := r.Registry[pfx]
	return registrySearchResult{
		Manufacturer: manuf,
		Found:        ok,
	}
}
