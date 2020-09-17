package imei_test

import (
	"testing"

	"github.com/alyyousuf7/imei-go"
)

func TestToDecimalIMEI(t *testing.T) {
	shouldPassCases := []string{"000000000000000000", "999999999999999999", "012345678901234567"}
	for _, str := range shouldPassCases {
		outIMEI, err := imei.DecimalIMEI.Parse(str)
		if err != nil {
			t.Errorf("Failed on parse %s: %s", str, err)
			continue
		}

		result := outIMEI.String(imei.DecimalIMEI)
		if result != str {
			t.Errorf("Expected %s, but got %s", str, result)
		}
	}
}

func TestToHexadecimalIMEI(t *testing.T) {
	shouldPassCases := []string{"00000000000000", "99999999999999"}
	for _, str := range shouldPassCases {
		outIMEI, err := imei.HexadecimalIMEI.Parse(str)
		if err != nil {
			t.Errorf("Failed on parse %s: %s", str, err)
			continue
		}

		result := outIMEI.String(imei.HexadecimalIMEI)
		if result != str {
			t.Errorf("Expected %s, but got %s", str, result)
		}
	}
}

func TestIMEICrossConversion(t *testing.T) {
	conversionTestCase := []struct {
		Decimal     string
		Hexadecimal string
		Checksum    string
	}{
		{"030541989609441844", "12345678901234", "7"},
		{"089379662400336177", "35464110052131", "2"},
		{"089555533708659600", "35611709842290", "2"},
		{"089798477702713159", "35862909296647", "6"},
		{"089594829606521636", "35671608638324", "8"},
	}

	for _, testcase := range conversionTestCase {
		// Decimal to Hexadecimal
		i, err := imei.DecimalIMEI.Parse(testcase.Decimal)
		if err != nil {
			t.Errorf("Failed on parse %s: %s", testcase.Decimal, err)
			continue
		}

		hex := i.String(imei.HexadecimalIMEI)
		if hex != testcase.Hexadecimal {
			t.Errorf("Expected Hexadecimal %s, but got %s", testcase.Hexadecimal, hex)
		}

		hex = i.String(imei.HexadecimalChecksumIMEI)
		if hex != testcase.Hexadecimal+testcase.Checksum {
			t.Errorf("Expected Hexadecimal with checksum %s, but got %s", testcase.Hexadecimal+testcase.Checksum, hex)
		}

		// Hexadecimal to Decimal
		i, err = imei.HexadecimalIMEI.Parse(testcase.Hexadecimal)
		if err != nil {
			t.Errorf("Failed on parse %s: %s", testcase.Hexadecimal, err)
			continue
		}

		dec := i.String(imei.DecimalIMEI)
		if dec != testcase.Decimal {
			t.Errorf("Expected Decimal %s, but got %s", testcase.Decimal, dec)
		}

		// Hexadecimal (Checksum) to Decimal
		i, err = imei.HexadecimalChecksumIMEI.Parse(testcase.Hexadecimal + testcase.Checksum)
		if err != nil {
			t.Errorf("Failed on parse %s: %s", testcase.Hexadecimal+testcase.Checksum, err)
			continue
		}

		dec = i.String(imei.DecimalIMEI)
		if dec != testcase.Decimal {
			t.Errorf("Expected Decimal %s, but got %s", testcase.Decimal, dec)
		}
	}
}
