package imei_test

import (
	"testing"

	"github.com/alyyousuf7/imei-go"
)

func TestGenerateChecksum(t *testing.T) {
	testcases := []struct {
		str      string
		checksum int
	}{
		{"00000000000000", 0},
		{"35464110052131", 2},
		{"35611709842290", 2},
		{"35862909296647", 6},
		{"35671608638324", 8},
	}

	for _, testcase := range testcases {
		checksum := imei.ChecksumFromString(testcase.str)
		if checksum != testcase.checksum {
			t.Errorf("Expected checksum %d for %s, but got %d", testcase.checksum, testcase.str, checksum)
		}
	}
}
