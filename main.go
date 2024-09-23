package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	var err error

	mode := flag.String("mode", "server", "server or client")

	url := flag.String("url", "localhost:8080", "url to connect to")

	flag.Parse()

	fmt.Println("Starting", *mode)

	switch *mode {
	case "server":
		err = startServer(*url)
	case "client":
		err = startClient(*url)
	default:
		err = fmt.Errorf("invalid mode: %s", *mode)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
