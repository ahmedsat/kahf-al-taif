package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/ahmedsat/bayaan"
	"github.com/ahmedsat/kahf-al-taif/client"
)

func init() {

	runtime.LockOSThread()

	// working directory is where we ar searching for files	like:
	// shaders, textures, etc.
	wd := os.Getenv("WORKING_DIRECTORY")
	if wd != "" {
		err := os.Chdir(wd)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	wd, _ = os.Getwd()

	// open log file
	f, err := os.OpenFile("../log.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		bayaan.Fatal(err.Error(), bayaan.Fields{
			"mode": "server",
			"url":  "localhost:8080",
		})
	}

	bayaan.Setup(
		bayaan.WithLevel(bayaan.LoggerLevelDebug),
		bayaan.WithTimeFormat("2006-01-02 15:04:05"),
		bayaan.WithOutput(f, false, false),
	)

	bayaan.Info("Working directory", bayaan.Fields{
		"workingDirectory": wd,
	})

}

func main() {

	var err error

	mode := flag.String("mode", "client", "server or client")
	url := flag.String("url", "localhost:8080", "url to connect to")

	bayaan.Debug("Parsing command line flags", bayaan.Fields{
		"mode": *mode,
		"url":  *url,
	})

	flag.Parse()

	switch *mode {
	case "server":
		err = startServer(*url)
	case "client":
		err = client.StartClient(*url)
	default:
		err = fmt.Errorf("invalid mode: %s", *mode)
	}

	if err != nil {
		bayaan.Fatal(err.Error(), bayaan.Fields{
			"mode": *mode,
			"url":  *url,
		})
	}
}
