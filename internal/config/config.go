package config

import (
	"errors"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	config *Config
)

type Config struct {
	Telegram *TelegramConfig `yaml:"telegram"`
}

type TelegramConfig struct {
	Token string `yaml:"token"`
}

func init() {
	if err := load(); err != nil {
		err := readFromEnv()
		if err != nil {
			panic(err)
		}
	}
}

func readFromEnv() error {
	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		return errors.New("token not set")
	}
	config = new(Config)
	config.Telegram = &TelegramConfig{
		Token: token,
	}
	return nil
}

func load() error {

	yamlFile, err := ioutil.ReadFile("configs/config.yml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return err
	}

	return nil
}

// Get returns the config or panics if not loaded.
func Get() *Config {
	return config
}
