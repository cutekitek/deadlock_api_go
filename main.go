package main

import (
	api "deadlock_analyzer/client"
	"deadlock_analyzer/proto/gc/deadlock"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/0xAozora/go-steam"
	"github.com/0xAozora/go-steam/protocol/protobuf"
	"gopkg.in/yaml.v3"
)



type ConsoleAuthentficator struct{}

func (c *ConsoleAuthentficator) GetCode(protobuf.EAuthSessionGuardType) string {
	var code string
	fmt.Print("Auth Code:")
	fmt.Scanln(&code)
	return code
}

type Config struct {
	Username string `yaml:"username"`
	AccessToken  string `yaml:"access_token"`
	RefreshToken string `yaml:"refresh_token"`
	GuardData    string `yaml:"guard_data"`
}

func loadConfig() (*Config, error) {
	configFile, err := os.Open("config.yml")
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	config := &Config{}
	err = yaml.NewDecoder(configFile).Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func main() {
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	details := &steam.LogOnDetails{
		Username: config.Username,
		AccessToken:  config.AccessToken,
		RefreshToken: config.RefreshToken,
		GuardData:    config.GuardData,
	}

	client := steam.NewClient()
	client.Auth.Authenticator = &ConsoleAuthentficator{}
	
	dlClient := api.NewHandler(client, details)
	if err := dlClient.Connect(); err != nil{
		panic(err)
	}
	steamId := client.SteamId().GetAccountId()
	dlClient.GetAccountStats(&deadlock.CMsgClientToGCGetAccountStats{
		AccountId: &steamId,
	})
	for evt := range dlClient.Events(){
		switch e := evt.(type) {
		case api.GetActiveMatchesResponse:
			fmt.Println("Active matches:", len(e.ActiveMatches))
			time.Sleep(time.Millisecond * 300)
			dlClient.GetActiveMatches(&deadlock.CMsgClientToGCGetActiveMatches{})

		case api.GetAccountStatsResponse:
			fmt.Println("Stats:", e.Stats.Stats)
		}
	}
}
