package imei

import (
	"fmt"
	"strconv"
)

var ErrInvalidIMEI = fmt.Errorf("invalid imei")

type IMEI struct {
	p1 int64
	p2 int64
}

func (imei IMEI) String(format IMEIFormat) string {
	if format.p1Width >= format.length {
		panic("invalid configuration")
	}

	p1 := strconv.FormatInt(imei.p1, format.base)
	p2 := strconv.FormatInt(imei.p2, format.base)

	out := fmt.Sprintf("%0*s%0*s", format.p1Width, p1, format.length-format.p1Width, p2)

	if format.checksum {
		checksum := ChecksumFromString(out)
		return fmt.Sprintf("%s%d", out, checksum)
	}

	return out
}

func (imei IMEI) Checksum() int {
	str := imei.String(HexadecimalIMEI)

	return ChecksumFromString(str)
}

type IMEIFormat struct {
	base     int
	length   int
	p1Width  int
	checksum bool
}

var DecimalIMEI = IMEIFormat{10, 18, 10, false}
var HexadecimalIMEI = IMEIFormat{16, 14, 8, false}
var HexadecimalChecksumIMEI = IMEIFormat{16, 14, 8, true}

func (format IMEIFormat) Parse(input string) (*IMEI, error) {
	var err error

	if (!format.checksum && len(input) != format.length) || (format.checksum && len(input)-1 != format.length) {
		return nil, fmt.Errorf("invalid length: %w", ErrInvalidIMEI)
	}

	if _, err = strconv.Atoi(input); err != nil {
		return nil, fmt.Errorf("invalid characters: %w", ErrInvalidIMEI)
	}

	imei := input
	var checksum int

	if format.checksum {
		imei = input[:len(imei)-1]

		checksum, err = strconv.Atoi(input[len(input)-1 : len(input)])
		if err != nil {
			return nil, fmt.Errorf("unable to parse checksum: %w", ErrInvalidIMEI)
		}

		if ChecksumFromString(imei) != checksum {
			return nil, fmt.Errorf("checksum mismatched: %w", ErrInvalidIMEI)
		}
	}

	p1, err := strconv.ParseInt(imei[:format.p1Width], format.base, 0)
	if err != nil {
		return nil, err
	}

	p2, err := strconv.ParseInt(imei[format.p1Width:], format.base, 0)
	if err != nil {
		return nil, err
	}

	_imei := IMEI{p1, p2}
	if _imei.String(format) != input {
		return nil, fmt.Errorf("unknown error occured, please contact developer with this info: %s: %w", _imei.String(format), ErrInvalidIMEI)
	}

	return &_imei, nil
}
