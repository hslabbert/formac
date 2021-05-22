# formac

`formac` formats MAC addresses into different common formats.

Install with `go get github.com/hslabbert/formac...`

MAC addresses are parsed using [`net.ParseMAC()`](https://golang.org/pkg/net/#ParseMAC), so any MAC address format supported by `net.ParseMAC()` is accepted by `formac`.

Output can be either `plain` or `json`, with the default being `plain`.

Outputted formats are:

- `Cisco`: Cisco format; three sets of 2 hex bytes separated by periods.
- `UnixExpanded`: What we commonly expect; hex bytes separated by colons.
- `UnixCompact`: Same as `Unix`, except duplicate zeroes removed.
- `EUI`: Dash-separated hex bytes, upper case.
- `Bare`: No separators, upper case.
- `PgSQL`: Two sets of 3 hex bytes, separated by a colon.

Also, will output the manufacturer ("OUI") for the MAC if found.  Manufacturer lookups are from [IEEE registries](https://standards.ieee.org/products-services/regauth/index.html).

## Usage

Basic usage:

```
$ formac -h
Usage of formac:
  -format string
        MAC output format (default "plain")
  string
        The MAC address to parse and format
```

Default to `plain` output:

```
$ formac 00-00-5E-00-53-01
Cisco: 0000.5e00.5301
UnixExpanded: 00:00:5e:00:53:01
UnixCompact: 0:0:5e:0:53:1
EUI: 00-00-5E-00-53-01
Bare: 00005E005301
PgSQL: 00005e:005301
Manufacturer: ICANN, IANA Department
```

Or as `json`:

```
$ formac -format json 0000.5e00.5301
{"Cisco":"0000.5e00.5301","UnixExpanded":"00:00:5e:00:53:01","UnixCompact":"0:0:5e:0:53:1","EUI":"00-00-5E-00-53-01","Bare":"00005E005301","PgSQL":"00005e:005301","Manufacturer":"IC
ANN, IANA Department"}
```
