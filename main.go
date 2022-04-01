package main

import (
    "log"
    "go.bug.st/serial"
)

func main() {
    ports, err := serial.GetPortsList()
    if err != nil {
        log.Fatal(err)
    }
}

