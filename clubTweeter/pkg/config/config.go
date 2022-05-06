package config

import (
	_ "embed"
	"log"

	"gopkg.in/yaml.v2"
)

//go:embed config.yaml
var ConfigString []byte

type TwitterAuth struct {
	ConsumerKey    string `yaml:"consumerKey"`
	ConsumerSecret string `yaml:"consumerSecret"`
	AccessToken    string `yaml:"accessToken"`
	AccessSecret   string `yaml:"accessSecret"`
}

type S3Path struct {
	Bucket string `yaml:"bucket"`
	Key    string `yaml:"key"`
}

type ScoreImgs struct {
	FiftyText string `yaml:"fiftyText"`
	HundredText string `yaml:"hundredText"`
	FiveWicketsText string `yaml:"fiveWktsText"`
	CachePath S3Path `yaml:"cachePath"`
}

type Config struct {
	PlayCricket struct {
		ClubID   string   `yaml:"clubID"`
		APIToken string   `yaml:"apiToken"`
		Teams    []string `yaml:"teams"`
	} `yaml:"playCricket"`
	PlayCricketTwitterMap  map[string]string `yaml:"playCricketTwitterMap"`
	TwitterAuth `yaml:"twitterAuth"`
	Templates   struct {
		Fixtures    string `yaml:"fixtures"`
		FixtureLine string `yaml:"fixtureLine"`
	} `yaml:"templates"`
	ScoreImgs `yaml:"scoreImgs"`
}

func NewConfig() Config {
	cfg := Config{}
	err := yaml.Unmarshal(ConfigString, &cfg)
	if err != nil {
		log.Fatalf("There was a problem loading config, %v", err)
	}
	return cfg
}
