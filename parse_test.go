package imei_test

import (
	"testing"

	"github.com/alyyousuf7/imei-go"
)

func TestParseDecimalIMEI(t *testing.T) {
	shouldFailCases := []string{"", "hello", "18 character text.", "FFFFFFFFFFFFFFFFFF", "ffffffffffffffffff"}
	for _, str := range shouldFailCases {
		_, err := imei.DecimalIMEI.Parse(str)
		if err == nil {
			t.Errorf("Expected to fail on input %s, but passed", str)
		}
	}

	shouldPassCases := []string{"000000000000000000", "999999999999999999", "012345678901234567"}
	for _, str := range shouldPassCases {
		_, err := imei.DecimalIMEI.Parse(str)
		if err != nil {
			t.Errorf("Failed on input %s: %s", str, err)
		}
	}
}

func TestParseHexadecialIMEI(t *testing.T) {
	shouldFailCases := []string{"", "hello", "14charactertex", "15charactertext", "abcdef", "abcdef0abcdefg", "FFFFFFFFFFFFFF", "ffffffffffffff"}
	for _, str := range shouldFailCases {
		_, err := imei.HexadecimalIMEI.Parse(str)
		if err == nil {
			t.Errorf("Expected to fail on input %s, but passed", str)
		}
	}

	shouldPassCases := []string{"00000000000000", "99999999999999"}
	for _, str := range shouldPassCases {
		_, err := imei.HexadecimalIMEI.Parse(str)
		if err != nil {
			t.Errorf("Failed on input %s: %s", str, err)
		}
	}
}
