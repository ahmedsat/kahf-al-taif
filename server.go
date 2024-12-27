package main

import (
	"fmt"

	"github.com/ahmedsat/bayaan"
	"github.com/ahmedsat/silah/global"
	"github.com/ahmedsat/silah/server"
)

type ClientMessageHandler struct {
}

func (mh *ClientMessageHandler) HandleMessage(clientID int, message global.Message) []global.Message {
	bayaan.Trace("Received message from client", bayaan.Fields{
		"clientID": clientID,
		"message":  message.Payload,
	})

	return []global.Message{
		global.NewMessage("client", fmt.Sprintf("Client %d: %s", clientID, message.Payload)),
	}
}

func (mh *ClientMessageHandler) ClientConnected(clientID int) []global.Message {
	bayaan.Info("Client Connected", bayaan.Fields{
		"clientID": clientID,
	})
	return []global.Message{
		global.NewMessage("client", fmt.Sprintf("Client %d Connected", clientID)),
	}
}

func (mh *ClientMessageHandler) ClientDisconnected(clientID int) []global.Message {
	bayaan.Info("Client Disconnected", bayaan.Fields{
		"clientID": clientID,
	})
	return []global.Message{
		global.NewMessage("client", fmt.Sprintf("Client %d Disconnected", clientID)),
	}
}

func startServer(url string) error {
	bayaan.Info("starting server", bayaan.Fields{
		"url": url,
	})
	cmh := ClientMessageHandler{}
	s := server.NewServer(&cmh)
	return s.Start(url)
}
