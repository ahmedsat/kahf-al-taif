package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ahmedsat/silah/client"
	"github.com/ahmedsat/silah/global"
	"github.com/ahmedsat/silah/server"
)

func main() {
	var err error

	mode := flag.String("mode", "server", "server or client")
	flag.Parse()

	fmt.Println("Starting", *mode)

	switch *mode {
	case "server":
		err = startServer()
	case "client":
		err = startClient()
	default:
		err = fmt.Errorf("invalid mode: %s", *mode)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type ClientMessageHandler struct {
}

func (mh *ClientMessageHandler) HandleMessage(clientID int, message global.Message) []global.Message {
	return []global.Message{
		global.NewMessage("client", fmt.Sprintf("Client %d: %s", clientID, message.Payload)),
	}
}

func (mh *ClientMessageHandler) ClientConnected(clientID int) []global.Message {
	return []global.Message{
		global.NewMessage("client", fmt.Sprintf("Client %d Connected", clientID)),
	}
}

func (mh *ClientMessageHandler) ClientDisconnected(clientID int) []global.Message {
	return []global.Message{
		global.NewMessage("client", fmt.Sprintf("Client %d Disconnected", clientID)),
	}
}

func startServer() error {
	cmh := ClientMessageHandler{}
	s := server.NewServer(&cmh)
	return s.Start(":8080")
}

func startClient() error {
	c, err := client.NewClient("localhost:8080")
	if err != nil {
		return err
	}

	go c.ReceiveMessages()

	ch := c.GetIncomingChannel()

	for message := range ch {
		fmt.Println(message)
	}

	return c.Close()
}
