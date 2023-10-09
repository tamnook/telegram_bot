package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	ConnString string `json:"conn_string"`
}

func New() *Config {
	return &Config{}
}

func (c *Config) ReadConfig() error {
	file, err := os.Open("tg_bot_resender/configs/config.json")
	if err != nil {
		return err
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}

	err = json.Unmarshal(fileBytes, c)
	if err != nil {
		return err
	}
	return nil
}
