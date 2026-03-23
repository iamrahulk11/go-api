package repository

import (
	"database/sql"
	"fmt"
	response "user-mapping/domain/dto/response/user"
	sqlwrapper "user-mapping/infrastructure"
)

type UserRepository struct {
	SQL *sqlwrapper.SQLWrapper
}

func NewUserRepository(sqlWrapper *sqlwrapper.SQLWrapper) *UserRepository {
	return &UserRepository{
		SQL: sqlWrapper,
	}
}

func (r *UserRepository) FetchAllUser() (*response.AllUserResponse, error) {
	return &response.AllUserResponse{
		Username: "rahul",
	}, nil
}

func (r *UserRepository) FetchUserProfile(User_id string) (*response.UserBasicDetailsResponse, error) {
	db, err := r.SQL.GetDB("sqlserver", "mt_infinity_conn")
	if err != nil {
		return nil, err
	}

	const query = `
		select employee_id username, employee_name name, employee_email email, employee_gender gender, is_active employement_status, employee_department department, employee_designation designation 
		from adm_user_details where employee_id = :user_id
	`
	var resp response.UserBasicDetailsResponse
	params := map[string]interface{}{
		"user_id": User_id,
	}

	rows, err := db.NamedQuery(query, params)
	if err != nil {
		return nil, fmt.Errorf("fetch user profile query failed: %w", err)
	}

	defer rows.Close()
	// handle no data
	if !rows.Next() {
		return nil, sql.ErrNoRows
	}
	// bind by column name
	if err := rows.StructScan(&resp); err != nil {
		return nil, fmt.Errorf("struct scan failed: %w", err)
	}

	return &resp, nil
}
