package config

import (
	_ "embed"
	"encoding/json"
	"log"
)

//go:embed config.json
var ConfigString []byte

type TwitterAuth struct {
	ConsumerKey    string `json:"consumerKey"`
	ConsumerSecret string `json:"consumerSecret"`
	AccessToken    string `json:"accessToken"`
	AccessSecret   string `json:"accessSecret"`
}

type Config struct {
	PlayCricket struct {
		ClubID   string   `json:"clubID"`
		APIToken string   `json:"apiToken"`
		Teams    []string `json:"teams"`
	} `json:"playCricket"`
	TwitterMap  map[string]string `json:"twitterMap"`
	TwitterAuth `json:"twitterAuth"`
	Templates   struct {
		Fixtures    string `json:"fixtures"`
		FixtureLine string `json:"fixtureLine"`
	} `json:"templates"`
}

func NewConfig() Config {
	cfg := Config{}
	err := json.Unmarshal(ConfigString, &cfg)
	if err != nil {
		log.Fatalf("There was a problem loading config, %v", err)
	}
	return cfg
}
