package main

import (
	"deadlock_analyzer/proto/gc/deadlock"
	"fmt"
	"log"
	"time"

	"github.com/0xAozora/go-steam"
	"github.com/0xAozora/go-steam/protocol"
	"github.com/0xAozora/go-steam/protocol/gamecoordinator"
	"github.com/0xAozora/go-steam/protocol/protobuf"
	"google.golang.org/protobuf/proto"
)

const AppId = 1422450

type handler struct{
	Client  *steam.Client 
	Details *steam.LogOnDetails
}

func (h *handler) HandleGCPacket(packet *gamecoordinator.GCPacket) {
	if packet.AppId != AppId {
		return
	}
	if msg, ok := deadlock.EGCCitadelClientMessages_name[int32(packet.MsgType)]; ok{
		fmt.Println("Deadlock client message:",msg)
		
	} else if msg, ok := deadlock.EGCCitadelServerMessages_name[int32(packet.MsgType)]; ok {
		fmt.Println("Deadlock server message:",msg)
	} else {
		fmt.Println("Unknown message id:", packet.MsgType)
	}
}

func (h *handler) HandlePacket(p *protocol.Packet) {
	fmt.Printf("Received packet: %+v\n", p)
}

func (h *handler) HandleDeadlockClientMsg(packet *gamecoordinator.GCPacket) {
	switch deadlock.EGCCitadelClientMessages(packet.MsgType) {
	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCGetAccountStatsResponse:
		stats := &deadlock.CMsgClientToGCGetAccountStatsResponse{}
		err := proto.Unmarshal(packet.Body, stats)
		if err != nil{
			panic(err)
		}
		fmt.Println("status", deadlock.CMsgClientToGCGetAccountStatsResponse_EResult_name[int32(*stats.Result)])

		if stats.Stats != nil{
			for _, h := range stats.Stats.Stats {
				fmt.Println(*h.HeroId, h.StatId, h.TotalValue)
			}
		}
	}
}

func (h *handler) HandleEvent(event interface{}) {
	switch e := event.(type) {
	case *steam.ConnectedEvent:
		h.Client.Auth.LogOn(h.Details)
	case *steam.LoggedOnEvent:
		log.Printf("Logged on (%v) with SteamId %v and account flags %v", e.Result, e.ClientSteamId, e.AccountFlags)
		h.Client.GC.SetGamesPlayed(AppId)
		time.Sleep(time.Millisecond*500)
		accId := h.Client.SteamId().GetAccountId()
		h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCGetAccountStats), &deadlock.CMsgClientToGCGetAccountStats{
			AccountId: &accId,
		}))
	case *steam.LoginKeyEvent:
		log.Printf("New LoginKey: %v\n", e.LoginKey)
	case *steam.LogOnFailedEvent:
		log.Printf("Failed to log on (%v)", e.Result)
	}

} 

type ConsoleAuthentficator struct {}

func (c *ConsoleAuthentficator) GetCode(protobuf.EAuthSessionGuardType) string {
	var code string
	fmt.Print("Auth Code:")
	fmt.Scanln(&code)
	return code
}

func main() {
	details := &steam.LogOnDetails{}
	client := steam.NewClient()
	client.Auth.Authenticator = &ConsoleAuthentficator{}
	h := &handler{Details: details, Client: client}
	client.RegisterPacketHandler(h)
	client.GC.RegisterPacketHandler(h)
	client.Connect()
	for evt := range client.Events() {
		h.HandleEvent(evt)
	}
}
