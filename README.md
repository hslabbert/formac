# formac

`formac` formats MAC addresses into different common formats.

Install with `go get github.com/hslabbert/formac...`

Example:

```
$ formac 00-00-5E-00-53-00
Cisco: 0000.5e00.5300
UnixExpanded: 00:00:5e:00:53:00
UnixCompact: 0:0:5e:0:53:0
EUI: 00-00-5E-00-53-00
Bare: 00005E005300
PgSQL: 00005e:005300
```
