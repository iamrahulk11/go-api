package sqlwrapper

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"
	"user-mapping/internal/config"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

// SQLWrapper holds connection strings and can return DB connections dynamically
type SQLWrapper struct {
	connConfigs map[string]string
	timeout     time.Duration
	dbs         map[string]*sqlx.DB
	mu          sync.Mutex
}

// NewSQLWrapper creates a wrapper from AppConfig
func NewSQLWrapper(cfg *config.AppConfig) (*SQLWrapper, error) {
	connConfigs := make(map[string]string)

	for name, c := range cfg.DBConfiguration.Connections {
		connConfigs[name] = c
	}

	return &SQLWrapper{
		connConfigs: connConfigs,
		timeout:     30 * time.Second,
		dbs:         make(map[string]*sqlx.DB),
	}, nil
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

// Example: Execute a query with parameters
func (s *SQLWrapper) ExecuteQuery(driver, name, sqlQuery string, params map[string]interface{}) (*sqlx.Rows, error) {
	db, err := s.GetDB(driver, name)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	stmt, err := db.PreparexContext(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var args []interface{}
	for k, v := range params {
		args = append(args, sql.Named(k, v))
	}

	rows, err := stmt.QueryxContext(ctx, args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
