package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug"`
	Listen  struct {
		Type   string `yaml:"type"`
		BindIp string `yaml:"bind_ip"`
		Port   string `yaml:"port"`
	} `yaml:"listen"`
	MongoDB struct {
		Host       string `yaml:"host"`
		Port       string `yaml:"port"`
		Username   string `yaml:"username"`
		Password   string `yaml:"password"`
		Database   string `yaml:"database"`
		AuthDB     string `yaml:"auth_db"`
		Collection string `yaml:"collection"`
	} `yaml:"mongodb"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		err := cleanenv.ReadConfig("config.yml", instance)
		if err != nil {
			desc, _ := cleanenv.GetDescription(instance, nil)
			fmt.Println(desc)
		}
	})
	return instance
}
