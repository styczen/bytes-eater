package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
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
	for {
		_, err := f.Read(b)
		if err != nil {
			log.Fatal(err)
		}

        // Convert byte value to HEX to pad with '0'
	    hexStr := strconv.FormatUint(uint64(b[0]), 16)
		if len(hexStr) == 1 {
			fmt.Printf("0%v ", hexStr)
		} else {
			fmt.Printf("%v ", hexStr)
		}

        // Check whether byte is "viewable" in a nice way
        if b[0] > 32 && b[0] < 127 {
			fmt.Printf("%v | ", string(b[0]))
        } else {
			fmt.Printf(". | ")
        }

        // Add new line if needed
        i += 1
        if i == *length {
            fmt.Println()
            i = 0
        }
	}
}
