package infrastructure

import (
	"context"
	"database/sql"
	"time"
	"user-mapping/internal/config"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

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
