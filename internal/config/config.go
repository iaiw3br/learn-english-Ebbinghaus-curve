package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Telegram Telegram `yaml:"telegram"`
	Client   struct {
		Postgresql `yaml:"postgresql"`
	} `yaml:"client"`
	Listen struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"listen"`
}

type Telegram struct {
	Token string `yaml:"token"`
}

type Postgresql struct {
	Driver   string `json:"driver"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var cfg *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		cfg = &Config{}
		if err := cleanenv.ReadConfig("config.yml", cfg); err != nil {
			log.Fatal(err)
		}
	})
	return cfg
}
