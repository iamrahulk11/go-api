package infrastructure

import (
	"context"
	"fmt"
	"sync"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// SQLWrapper holds connection strings and can return DB connections dynamically
type SQLWrapper struct {
	connConfigs map[string]string
	timeout     time.Duration
	dbs         map[string]*sqlx.DB
	mu          sync.Mutex
}

// GetDB returns a *sqlx.DB for a connection name
func (s *SQLWrapper) GetDB(driver string, name string) (*sqlx.DB, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// return cached DB
	if db, ok := s.dbs[name]; ok {
		return db, nil
	}

	conn, ok := s.connConfigs[name]
	if !ok {
		return nil, fmt.Errorf("connection name %s not found", name)
	}

	db, err := sqlx.Connect(driver, conn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
