package config

import (
	"encoding/json"
	"fmt"
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
	AppEnv          string          `json:"APP_ENV"`
	Port            string          `json:"PORT"`
	JWT             JWTConfig       `json:"jwtConfiguration"`
	DBConfiguration DBConfiguration `json:"DBConfiguration"`
}

func LoadConfig(baseFilePath string) (*AppConfig, error) {
	// Load base config
	config, err := loadFile(baseFilePath)
	if err != nil {
		return nil, err
	}

	// Load environment-specific config if APP_ENV is set
	if config.AppEnv != "" {
		envFilePath := fmt.Sprintf("appsettings.%s.json", config.AppEnv)
		if _, err := os.Stat(envFilePath); err == nil {
			envConfig, err := loadFile(envFilePath)
			if err != nil {
				return nil, err
			}
			config = envConfig // just replace it entirely
		}
	}

	return config, nil
}

// helper function to read JSON file into AppConfig
func loadFile(filePath string) (*AppConfig, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file %s: %w", filePath, err)
	}
	defer file.Close()

	var config AppConfig
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode config file %s: %w", filePath, err)
	}

	return &config, nil
}
