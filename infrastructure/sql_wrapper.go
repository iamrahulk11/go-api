package infrastructure

import (
	"context"
	"database/sql"
	"user-mapping/internal/config"

	_ "github.com/denisenkom/go-mssqldb"
)

type SQLWrapper struct {
	manager *config.DBConnections
}

func NewSQLWrapper(m *config.DBConnections) *SQLWrapper {
	return &SQLWrapper{
		manager: m,
	}
}

// Example: Execute a query with parameters
func (s *SQLWrapper) ExecuteQuery(connName string, sqlQuery string, params map[string]interface{}) ([]map[string]interface{}, error) {
	connString := s.manager.Connections[connName].ConnectionString
	timeout := s.manager.Connections[connName].Timeout

	dbM := GetDBManager()
	db, err := dbM.OpenDB("mssql", connString, timeout)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5) // max time to connect
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

	cols, _ := rows.Columns()

	var results []map[string]interface{}

	for rows.Next() {
		values := make([]interface{}, len(cols))
		ptrs := make([]interface{}, len(cols))

		for i := range values {
			ptrs[i] = &values[i]
		}

		if err := rows.Scan(ptrs...); err != nil {
			return nil, err
		}

		row := make(map[string]interface{})
		for i, col := range cols {
			if b, ok := values[i].([]byte); ok {
				row[col] = string(b)
			} else {
				row[col] = values[i]
			}
		}

		results = append(results, row)
	}

	return results, rows.Err()
}
