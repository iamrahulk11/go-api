package config

import (
	"fmt"
	"time"
)

type DBConnConfig struct {
	Driver  string
	DSN     string
	Timeout time.Duration
}

type DBConfig struct {
	Connections map[string]DBConnConfig
}

// DatabaseConnectionConfig creates a wrapper from AppConfig
func NewDBConfig(connections map[string]DBConnConfig) *DBConfig {
	return &DBConfig{
		Connections: connections,
	}
}

// GetConnectionString returns the connection string for a given DB name
func (d *DBConfig) GetConnectionString(name string) (string, error) {
	conn, exists := d.Connections[name]
	if !exists {
		return "", fmt.Errorf("connection config not found for: %s", name)
	}

	return conn.DSN, nil
}

func (d *DBConfig) Timeout(name string) (time.Duration, error) {
	conn, exists := d.Connections[name]
	if !exists {
		return 0, fmt.Errorf("connection config not found for: %s", name)
	}
	return conn.Timeout, nil
}
