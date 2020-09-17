package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/alyyousuf7/imei-go"
)

func main() {
	// Flags
	allowedInputFormats := []string{"auto", "hex-checksum", "hex", "dec"}
	allowedOutputFormats := []string{"hex-checksum", "hex", "dec"}

	var inputFormatStr string
	flag.StringVar(&inputFormatStr, "input-format", allowedInputFormats[0], fmt.Sprintf("input imei `format` [%s]", strings.Join(allowedInputFormats, ", ")))

	var outputFormatStr string
	flag.StringVar(&outputFormatStr, "output-format", allowedOutputFormats[0], fmt.Sprintf("output imei `format` [%s]", strings.Join(allowedOutputFormats, ", ")))

	var inputFile string
	flag.StringVar(&inputFile, "input", "-", "input `path` (use - for stdin)")

	flag.Parse()

	var outputFormat imei.IMEIFormat
	switch outputFormatStr {
	case "dec":
		outputFormat = imei.DecimalIMEI
	case "hex":
		outputFormat = imei.HexadecimalIMEI
	case "hex-checksum":
		outputFormat = imei.HexadecimalChecksumIMEI
	default:
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Read
	var file io.Reader

	if inputFile == "-" {
		file = os.Stdin
	} else {
		file, err := os.Open(inputFile)
		if err != nil {
			fmt.Println(err)
			fmt.Println()

			fmt.Println("Usage:")
			flag.PrintDefaults()
			os.Exit(1)
		}
		defer file.Close()
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		var parsedIMEI *imei.IMEI
		var err error

		switch inputFormatStr {
		case "auto":
			parsedIMEI, err = imei.DecimalIMEI.Parse(text)
			if err != nil {
				parsedIMEI, err = imei.HexadecimalIMEI.Parse(text)
				if err != nil {
					parsedIMEI, err = imei.HexadecimalChecksumIMEI.Parse(text)
				}
			}
		case "dec":
			parsedIMEI, err = imei.DecimalIMEI.Parse(text)
		case "hex":
			parsedIMEI, err = imei.HexadecimalIMEI.Parse(text)
		case "hex-checksum":
			parsedIMEI, err = imei.HexadecimalChecksumIMEI.Parse(text)
		default:
			fmt.Println("unknown input format")
			fmt.Println()

			fmt.Println("Usage:")
			flag.PrintDefaults()
			os.Exit(1)
			break
		}

		if err != nil {
			fmt.Println(text, "-", err)
		} else if parsedIMEI != nil {
			fmt.Println(parsedIMEI.String(outputFormat))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
