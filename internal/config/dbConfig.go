package config

import (
	"time"
)

type DBConn struct {
	ConnectionString string
	Timeout          time.Duration
}

type DBConnections struct {
	Connections map[string]DBConn
}

func (c DBConfiguration) MapDBConnections() *DBConnections {
	dbConnections := make(map[string]DBConn)

	for key, connStr := range c.Connections {
		dbConnections[key] = DBConn{
			ConnectionString: connStr,
			Timeout:          30 * time.Second,
		}
	}

	return &DBConnections{
		Connections: dbConnections,
	}
}
