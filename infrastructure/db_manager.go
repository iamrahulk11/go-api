package infrastructure

import (
	"fmt"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
)

type DBManager struct {
	mu    sync.RWMutex
	pools map[string]*sqlx.DB
}

var (
	instance *DBManager
	once     sync.Once
)

func GetDBManager() *DBManager {
	once.Do(func() {
		instance = &DBManager{
			pools: make(map[string]*sqlx.DB),
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
func (m *DBManager) OpenDB(driver, ConnectionString string, Timeout time.Duration) (*sqlx.DB, error) {

	// fast path (read lock)
	m.mu.RLock()
	db, ok := m.pools[ConnectionString]
	m.mu.RUnlock()

	if ok {
		return db, nil
	}

	// write lock
	m.mu.Lock()
	defer m.mu.Unlock()

	// double-check
	if db, ok := m.pools[ConnectionString]; ok {
		return db, nil
	}

	conn, err := sqlx.Open(driver, ConnectionString)
	if err != nil {
		return nil, err
	}

	// verify connection
	if err := conn.Ping(); err != nil {
		return nil, err
	}

	// timeout handling
	if Timeout > 0 {
		conn.SetConnMaxLifetime(Timeout)
	} else {
		conn.SetConnMaxLifetime(30 * time.Second)
	}

	m.pools[ConnectionString] = conn
	return conn, nil
}
