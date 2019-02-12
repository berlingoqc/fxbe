package api

import (
	"encoding/json"
	"io/ioutil"
)

var Version string

// Config is the structure for the configuration of the api
type Config struct {
	Bind string `json:"bind"`
	Root string `json:"root"`
}

// Save ...
func (c *Config) Save(filepath string) error {
	data, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath, data, 0644)
}

// LoadConfig ...
func LoadConfig(filepath string) (*Config, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	c := &Config{}
	return c, json.Unmarshal(data, c)
}
