package config

import (
	_ "embed"
	"encoding/json"
	"log"
)

//go:embed config.json
var ConfigString []byte

type Config struct {
	ClubID      string            `json:"clubID"`
	APIToken    string            `json:"apiToken"`
	Teams       []string          `json:"teams"`
	TwitterMap  map[string]string `json:"twitterMap"`
	TwitterAuth struct {
		ConsumerKey    string `json:"consumerKey"`
		ConsumerSecret string `json:"consumerSecret"`
		AccessToken    string `json:"accessToken"`
		AccessSecret   string `json:"accessSecret"`
	} `json:"twitterAuth"`
	Templates struct {
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
