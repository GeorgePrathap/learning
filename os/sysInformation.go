package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
)

func main() {
	currentDirectory, err := os.Getwd()
	if err != nil {
		fmt.Printf("error in getting the current directory: %v\n", err.Error())
		return
	}

	fmt.Printf("current directory: %v\n", currentDirectory)

	hostName, err := os.Hostname()
	if err != nil {
		fmt.Printf("error in getting the host name: %v\n", err)
		return
	}

	fmt.Printf("host name: %v\n", hostName)

	numOfCpu := runtime.NumCPU()

	fmt.Printf("number of cpu count: %v\n", numOfCpu)

	fmt.Printf("os: %v\n", runtime.GOOS)

	fmt.Printf("architecture: %v\n", runtime.GOARCH)

	// go findOpenPort()
}

func findOpenPort() {
	for port := 1; port < 65535; port++ {
		address := fmt.Sprintf("localhost:%d", port)
		connection, err := net.Dial("tcp", address)
		if err != nil {
			continue
		}
		connection.Close()

		fmt.Printf("open port: %v\n", port)
	}
}
