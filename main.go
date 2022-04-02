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
	flag.Parse()

	fileInfo, err := os.Stat(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	if fileInfo.IsDir() {
		log.Fatal("passed path is directory and it should be a file")
	}

	f, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b := make([]byte, 1)
	for {
		_, err := f.Read(b)
		if err != nil {
			log.Fatal(err)
		}
		hexStr := strconv.FormatUint(uint64(b[0]), 16)
		if len(hexStr) == 1 {
			fmt.Printf("0%v ", hexStr)
		} else {
			fmt.Printf("%v ", hexStr)
		}
	}
}
