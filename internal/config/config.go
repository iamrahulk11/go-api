package config

import (
	"encoding/json"
	"os"
)

type DBConfiguration struct {
	Connections map[string]string `json:"Connections"`
}

type JWTConfig struct {
	Issuer          string `json:"Issuer"`
	Audience        string `json:"Audience"`
	Secret          string `json:"Secret"`
	ExpiresInMinute int    `json:"ExpiresInMinute"`
}

type AppConfig struct {
	JWT             JWTConfig       `json:"jwtConfiguration"`
	DBConfiguration DBConfiguration `json:"DBConfiguration"`
}

func LoadConfig(filePath string) (*AppConfig, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config AppConfig
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
