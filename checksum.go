package imei

import (
	"strconv"
	"strings"
)

func ChecksumFromString(str string) int {
	sum := 0
	for i, ch := range strings.Split(str, "") {
		n, _ := strconv.Atoi(ch)

		if (i % 2) == 1 {
			n = n * 2
		}

		if n >= 10 {
			n = (n % 10) + 1
		}

		sum = sum + n
	}

	return (10 - (sum % 10)) % 10
}
