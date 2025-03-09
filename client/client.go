package api

import (
	"deadlock_analyzer/proto/gc/deadlock"
	"fmt"
	"log"
	"time"

	"github.com/0xAozora/go-steam"
	"github.com/0xAozora/go-steam/protocol"
	"github.com/0xAozora/go-steam/protocol/gamecoordinator"
)

const AppId = 1422450

type Client struct {
	Client  *steam.Client
	Details *steam.LogOnDetails
	events chan interface{}
}

func NewHandler(client *steam.Client, details *steam.LogOnDetails) *Client {
	cl := &Client{
		Client:  client,
		Details: details,
		events: make(chan interface{}, 3),
	}
	client.RegisterPacketHandler(cl)
	client.GC.RegisterPacketHandler(cl)
	return cl
}

func (h *Client) HandleGCPacket(packet *gamecoordinator.GCPacket) {
	if packet.AppId != AppId {
		return
	}
	if msg, ok := deadlock.EGCCitadelClientMessages_name[int32(packet.MsgType)]; ok {
		fmt.Println("Deadlock client message:", msg)
		response, err := h.ParseGCResponse(packet)
		if err != nil {
			log.Println("failed to read response:", err)
			return
		}
		h.events <- response
	} else if msg, ok := deadlock.EGCCitadelServerMessages_name[int32(packet.MsgType)]; ok {
		fmt.Println("Deadlock server message:", msg)
	} else {
		fmt.Println("Unknown message id:", packet.MsgType)
	}
}

func (h *Client) Connect() (error) {
	port, err := h.Client.Connect()
	if err != nil{
		return err
	}
	fmt.Println("Connected to:", port.String())

	for evt := range h.Client.Events() {
		switch e := evt.(type) {
		case *steam.ConnectedEvent:
			h.Client.Auth.LogOn(h.Details)
		case *steam.LoggedOnEvent:
			log.Printf("Logged on (%v) with SteamId %v and account flags %v", e.Result, e.ClientSteamId, e.AccountFlags)
			h.Client.GC.SetGamesPlayed(AppId)
			time.Sleep(time.Millisecond * 500)
			return nil
		case *steam.LogOnFailedEvent:
			return fmt.Errorf("log on failed: %s", e.Result.String())
		}
	}
	return fmt.Errorf("log on failed: no events to process")
}

func (h *Client) HandlePacket(p *protocol.Packet) {
	fmt.Printf("Received packet: %+v\n", p)
}

func (h *Client) Events() <-chan interface{}{
	return h.events
}