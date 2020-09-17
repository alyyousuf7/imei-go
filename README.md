# IMEI converter
This tool takes an IMEI in one format and outputs in another.

## Supported IMEI formats
- 18 digit decimal
- 14 digit hexadecimal
- 15 digit hexadecimal with checksum

## Build
```sh
$ make build
$ ./imei -h
Usage of ./imei:
  -input path
        input path (use - for stdin) (default "-")
  -input-format format
        input imei format [auto, hex-checksum, hex, dec] (default "auto")
  -output-format format
        output imei format [hex-checksum, hex, dec] (default "hex-checksum")
```