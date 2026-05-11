package infrastructure

import (
	"fmt"
	"sync"
	"time"
	"user-mapping/internal/config"

	"github.com/jmoiron/sqlx"
)

type DBManager struct {
	mu sync.RWMutex

	configs map[string]config.DBConnConfig
	pools   map[string]*sqlx.DB
}

var (
	instance *DBManager
	once     sync.Once
)

func GetDBManager(cfg *config.DBConfig) *DBManager {
	once.Do(func() {
		instance = &DBManager{
			configs: cfg.Connections,
			pools:   make(map[string]*sqlx.DB),
		}
	})
	return instance
}

func (m *DBManager) CloseAll() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for name, db := range m.pools {
		if err := db.Close(); err != nil {
			return fmt.Errorf("error closing %s: %w", name, err)
		}
		delete(m.pools, name)
	}

	return nil
}
func (m *DBManager) OpenDB(driver, name string) (*sqlx.DB, error) {

	// fast path (read lock)
	m.mu.RLock()
	db, ok := m.pools[name]
	m.mu.RUnlock()

	if ok {
		return db, nil
	}

	// write lock
	m.mu.Lock()
	defer m.mu.Unlock()

	// double-check
	if db, ok := m.pools[name]; ok {
		return db, nil
	}

	cfg, exists := m.configs[name]
	if !exists {
		return nil, fmt.Errorf("db config not found: %s", name)
	}

	conn, err := sqlx.Open(driver, name)
	if err != nil {
		return nil, err
	}

	// verify connection
	if err := conn.Ping(); err != nil {
		return nil, err
	}

	// timeout handling
	if cfg.Timeout > 0 {
		conn.SetConnMaxLifetime(cfg.Timeout)
	} else {
		conn.SetConnMaxLifetime(30 * time.Second)
	}

	m.pools[name] = conn
	return conn, nil
}
