package config

import (
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

const (
	_defaultCfgFilePath = "config/config.yaml"

	_defaultServerPot = "3000"
)

var (
	once sync.Once
	cfg  *Config
)

type Config struct {
	Slack  SlackData  `yaml:"slack"`
	Server ServerData `yaml:"server"`
}

type ServerData struct {
	Port string `yaml:"port"`
}

type SlackData struct {
	BotToken      string `yaml:"botToken"`
	SigningSecret string `yaml:"signingSecret"`
}

func InitCfg(path string) {
	var initialCfg Config
	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	if err := yaml.Unmarshal(b, &initialCfg); err != nil {
		log.Fatal(err)
	}

	if initialCfg.Server.Port == "" {
		initialCfg.Server.Port = _defaultServerPot
	}

	cfg = &initialCfg
}

func Get() *Config {
	if cfg == nil {
		once.Do(func() {
			InitCfg(_defaultCfgFilePath)
		})
	}
	return cfg
}
