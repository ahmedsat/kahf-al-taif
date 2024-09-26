package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/ahmedsat/bayaan"
)

func init() {
	runtime.LockOSThread()

	// working directory is where we ar searching for files	like:
	// shaders, textures, etc.
	wd := os.Getenv("WORKING_DIRECTORY")
	fmt.Println(wd)
	if wd != "" {
		err := os.Chdir(wd)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println(os.Getwd())

}

func main() {

	// bayaan.SetLevel(bayaan.LoggerLevelTrace)

	var err error
	mode := flag.String("mode", "client", "server or client")
	url := flag.String("url", "localhost:8080", "url to connect to")

	bayaan.Trace("Parsing command line flags")
	flag.Parse()

	switch *mode {
	case "server":
		err = startServer(*url)
	case "client":
		err = startClient(*url)
	default:
		err = fmt.Errorf("invalid mode: %s", *mode)
	}

	if err != nil {
		bayaan.Fatal("%s", err)
	}
}
