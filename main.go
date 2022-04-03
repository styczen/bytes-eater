package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileName := flag.String("f", "", "file to read from")
	length := flag.Uint("l", 16, "number of bytes displayed in one line")
	flag.Parse()

	// Check whether passed path is path to file
	fileInfo, err := os.Stat(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	if fileInfo.IsDir() {
		log.Fatal("passed path is a directory but it should be a path to file")
	}

	// Open file and handle error
	f, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b := make([]byte, 1)
	var i uint
	var bytesStr, charsStr string
	for {
		n, err := f.Read(b)
		if n == 0 {
			fmt.Println()
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// Convert byte value to HEX to pad with '0'
		hexStr := strconv.FormatUint(uint64(b[0]), 16)
		if len(hexStr) == 1 {
			bytesStr += fmt.Sprintf("0%v ", hexStr)
		} else {
			bytesStr += fmt.Sprintf("%v ", hexStr)
		}

		nrSpaces := *length - (i + 1)
		bytesWithSpaces := bytesStr + strings.Repeat("   ", int(nrSpaces))
		bytesWithSpaces += "| "

		// Check whether byte is "viewable" in a nice way
		if b[0] > 32 && b[0] < 127 {
			charsStr += fmt.Sprintf("%v", string(b[0]))
		} else {
			charsStr += "."
		}

		// Print prepared line
		fmt.Printf("\r%v", bytesWithSpaces+charsStr)

		// Add new line if needed
		i += 1
		if i == *length {
			fmt.Println()
			i = 0
			bytesStr, charsStr = "", ""
		}
	}
}
